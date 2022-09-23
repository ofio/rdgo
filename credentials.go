package rdgo

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"
)

func GetServiceAccountCredentials() (*google.Credentials, error) {
	var credentials *google.Credentials
	scopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/idtoken",
	}

	json, _ := ioutil.ReadFile("./service_account.json")
	var err error
	ctx := context.Background()
	if len(json) > 0 {
		credentials, err = google.CredentialsFromJSON(ctx, json, scopes...)
		if err != nil {
			log.Println("could not find credentials")
			return nil, err
		}
	} else {
		credentials, err = google.FindDefaultCredentials(ctx, scopes...)
		if err != nil {
			log.Println("could not find default credentials")
			return nil, err
		}
	}
	return credentials, nil
}

type CredentialsStruct struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func GenerateServiceAccountJWT(expirationSec int64, aud string) (string, error) {
	credentials, err := GetServiceAccountCredentials()
	if err != nil {
		log.Println(err)
		return "", err
	}

	var credentialsClone CredentialsStruct
	err = json.Unmarshal(credentials.JSON, &credentialsClone)
	if err != nil {
		log.Println(err)
		return "", err
	}

	credentialsCloneBytes, err := json.Marshal(credentialsClone)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if credentialsClone.Type == "authorized_user" {
		return "", &MyError{"credential type authorized_user not allowed. service account required"}
	}

	gcpScope := "https://www.googleapis.com/auth/cloud-platform"
	idScope := "https://www.googleapis.com/auth/idtoken"
	conf, err := google.JWTConfigFromJSON(credentialsCloneBytes, gcpScope, idScope)
	if err != nil {
		log.Println(err)
		return "", err
	}

	now := time.Now().Unix()

	// Build the JWT payload.
	jwtPayload := &jws.ClaimSet{
		Iat: now,
		// expires after 'expiryLength' seconds.
		Exp: now + expirationSec,
		// Iss must match 'issuer' in the security configuration in your
		// swagger spec (e.g. service account email). It can be any string.
		Iss: conf.Email,
		// Aud must be either your Endpoints service name, or match the value
		// specified as the 'x-google-audience' in the OpenAPI document.
		Aud: aud,
		// Sub and Email should match the service account's email address.
		Sub:           conf.Email,
		PrivateClaims: map[string]interface{}{"email": conf.Email},
	}
	jwsHeader := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
		KeyID:     conf.PrivateKeyID,
	}
	block, _ := pem.Decode(conf.PrivateKey)
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("private key parse error: %v", err)
	}
	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	// Sign the JWT with the service account's private key.
	if !ok {
		return "", errors.New("private key failed rsa.PrivateKey type assertion")
	}
	return jws.Encode(jwsHeader, jwtPayload, rsaKey)
}

type JwtStruct struct {
	Header    string
	Payload   string
	Signature string
}

func ParseJWT(idToken string) (*JwtStruct, error) {
	segments := strings.Split(idToken, ".")
	if len(segments) != 3 {
		return nil, fmt.Errorf("idtoken: invalid token, token must have three segments; found %d", len(segments))
	}
	return &JwtStruct{
		Header:    segments[0],
		Payload:   segments[1],
		Signature: segments[2],
	}, nil
}

func DecodeB64String(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}

// jwtHeader represents a parted jwt's header segment.
type JwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	KeyID     string `json:"kid"`
}
