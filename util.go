package rdgo

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/storage"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//verifyToken verifies jws signature
func VerifyToken(token string, rootPEM string) (bool, *rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(rootPEM))
	if block == nil || block.Type != "CERTIFICATE" {
		log.Fatal("failed to decode PEM block containing public key")
	}
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)

	pub := cert.PublicKey.(*rsa.PublicKey)
	parts := strings.Split(token, ".")
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], pub)
	if err != nil {
		return false, pub, nil
	}
	return true, pub, nil
}

func ReadObj(object string, iid string, bucket string) (*storage.ObjectAttrs, []byte, error) {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	// [START download_file]

	rc, err := client.Bucket(bucket).Object(iid + "/" + object).NewReader(ctx)

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	bh := client.Bucket(bucket)

	obj := bh.Object(iid + "/" + object)
	attrs, err := obj.Attrs(ctx)

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, rc)
	fb := buf.Bytes()
	return attrs, fb, err
}

// accessSecretVersion accesses the payload for the given secret version if one
// exists. The version can be a version number as a string (e.g. "5") or an
// alias (e.g. "latest").
func AccessSecretVersion(name string) (string, error) {
	// name := "projects/my-project/secrets/my-secret/versions/5"
	// name := "projects/my-project/secrets/my-secret/versions/latest"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}
func StandardAuth(r *http.Request, rootPEM string) (string, int, string, string, error) {
	var uid string = ""
	var iid int = -1
	var tkn string = ""
	var email string = ""

	c := r.Header.Get("Authorization")
	if len(c) < 10 {
		return tkn, iid, uid, email, &MyError{"authorization header missing"}
	} else {
		// Get the JWT string from the cookie
		tkn = strings.Replace(c, `Bearer `, "", -1)
	}

	isValid, pub, err := VerifyToken(tkn, rootPEM)
	if err != nil {
		return tkn, iid, uid, email, &MyError{"cannot verify token"}
	}

	if !isValid {
		return tkn, iid, uid, email, &MyError{"invalid token"}
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tkn, claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		return tkn, iid, uid, email, err
	}

	num, claimsOk := claims["https://hasura.io/jwt/claims"]
	if !claimsOk {
		return tkn, iid, uid, email, &MyError{"claims key missing"}
	}

	md, interfaceOk := num.(map[string]interface{})
	if !interfaceOk {
		return tkn, iid, uid, email, &MyError{"cannot type cast claims"}
	}

	uidstr, uidOk := md["x-hasura-user-id"]
	if !uidOk {
		return tkn, iid, uid, email, &MyError{"user id missing from claims"}
	} else {
		uid = uidstr.(string)
	}

	emailStr, emailStrOk := md["x-hasura-email"]
	if !emailStrOk {
		return tkn, iid, uid, email, &MyError{"email missing from claims"}
	}
	email = emailStr.(string)

	iidStr, iidStrOk := md["x-hasura-instance-id"]
	if !iidStrOk {
		return tkn, iid, uid, email, &MyError{"instance id missing from claims"}
	}

	iid, err = strconv.Atoi(iidStr.(string))
	if err != nil {
		return "", -1, "", email, &MyError{"instance id missing from claims"}
	}

	return tkn, iid, uid, email, nil
}

func Objcheck(bucket, name string) bool {
	client, err := google.DefaultClient(oauth2.NoContext,
		"https://www.googleapis.com/auth/devstorage.full_control")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/storage/v1/b/"+bucket+"/"+"o/"+url.QueryEscape(name), nil)
	//if the object doesnt exist get returns a 404 status error
	resp, err := client.Do(req)
	var exist bool
	exist = true
	log.Println(resp.StatusCode)
	if resp.StatusCode == 404 {
		exist = false
	}

	if err != nil {
		exist = false
	}
	return exist
}

func GcsUpload(ctx context.Context, client *storage.Client, r bufio.Reader, bucket string, name string, mime string, requestfilename string) error {
	bh := client.Bucket(bucket)
	// Next check if the bucket exists
	if _, err := bh.Attrs(ctx); err != nil {
		return err
	}

	obj := bh.Object(name)
	pr, pw := io.Pipe()

	go func() {

		gw := gzip.NewWriter(pw)

		_, err := r.WriteTo(gw)
		if err != nil {
			fmt.Println(err)
		}
		gw.Close()
		pw.Close()
	}()
	ow := obj.NewWriter(ctx)
	ow.ContentEncoding = "gzip"

	ow.ContentType = mime
	Metadata := make(map[string]string)
	if requestfilename != "" {
		Metadata["filename"] = requestfilename
	}
	ow.Metadata = Metadata
	if _, err := io.Copy(ow, pr); err != nil {
		return err
	}

	defer ow.Close()
	/*
		if public {
			if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
				return nil, nil, err
			}
		}
	*/

	return nil
}

func MustGetEnv(envKey, defaultValue string) string {
	val := os.Getenv(envKey)
	if val == "" {
		val = defaultValue
	}
	if val == "" {
		log.Fatalf("%q should be set", envKey)
	}
	return val
}

func UpsertAttachmentGen(tableName string, attachmentName string, attachmentUuid string, attachmentGen int64, attachmentMime string, attachmentInstance int, attachmentUser string, objectId int, endpoint string, adminSecret string, bearer string, createdByInstanceID int, updatedByInstanceID int, constraint string) (int, string, error) {
	upsertAttachmentGQL := `mutation upsert_` + tableName + `_attachment($changes: [` + tableName + `_attachment_insert_input!]!) {
		insert_` + tableName + `_attachment(objects: $changes, on_conflict: {constraint: ` + constraint + `, update_columns: [name, uuid, generation, mime_type, read_secret]}) {
			affected_rows
			returning {
				id
				uuid
			}
		}
	}	
`
	if len(attachmentMime) < 1 {
		attachmentMime = "application/octet-stream"
	}
	changes := map[string]interface{}{
		"name":        attachmentName,
		"uuid":        attachmentUuid,
		"generation":  attachmentGen,
		"mime_type":   attachmentMime,
		"instance_id": attachmentInstance,
		"created_by":  attachmentUser,
		"updated_by":  attachmentUser,
	}

	changes[tableName+"_id"] = objectId
	if createdByInstanceID > -1 {
		changes["created_by_instance_id"] = createdByInstanceID
	}
	if updatedByInstanceID > -1 {
		changes["updated_by_instance_id"] = updatedByInstanceID
	}

	payload := map[string]interface{}{
		"query": upsertAttachmentGQL,
		"variables": map[string]interface{}{
			"changes": changes,
		},
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)

	reqAPI, _ := http.NewRequest("POST", endpoint, b)
	log.Print(reqAPI.Body)

	reqAPI.Header.Add("content-type", "application/json")
	if len(adminSecret) > 0 {
		reqAPI.Header.Add("X-Hasura-Admin-Secret", adminSecret)
	} else {
		reqAPI.Header.Add("Authorization", bearer)
	}

	resAPI, err := http.DefaultClient.Do(reqAPI)
	rd := Responsedata{}
	if err != nil {
		fmt.Println("hasura api error: ", err)
		return -1, "", err
	}

	decoder := json.NewDecoder(resAPI.Body)
	err = decoder.Decode(&rd)
	if err != nil {
		fmt.Println(err)
		return -1, "", err
	}

	if len(rd.Errors) != 0 {
		log.Println("hasura error: ", rd.Errors)
		return -1, "", &MyError{rd.Errors[0].Message}
	} else {
		if len(rd.Data.InsertContractAttachment.Returning) > 0 {
			return rd.Data.InsertContractAttachment.Returning[0].Id, rd.Data.InsertContractAttachment.Returning[0].UUID, nil
		} else if len(rd.Data.InsertPoHeaderAttachment.Returning) > 0 {
			return rd.Data.InsertPoHeaderAttachment.Returning[0].Id, rd.Data.InsertPoHeaderAttachment.Returning[0].UUID, nil
		} else if len(rd.Data.InsertInvoiceAttachment.Returning) > 0 {
			return rd.Data.InsertInvoiceAttachment.Returning[0].Id, rd.Data.InsertInvoiceAttachment.Returning[0].UUID, nil
		} else if len(rd.Data.ContractAttachment) > 0 {
			return rd.Data.ContractAttachment[0].ID, rd.Data.ContractAttachment[0].UUID, nil
		} else if len(rd.Data.PoHeaderAttachment) > 0 {
			return rd.Data.PoHeaderAttachment[0].ID, rd.Data.PoHeaderAttachment[0].UUID, nil
		} else if len(rd.Data.InvoiceAttachment) > 0 {
			return rd.Data.InvoiceAttachment[0].ID, rd.Data.InvoiceAttachment[0].UUID, nil
		}
	}
	return -1, "", &MyError{"failed to upsert attachment"}
}

func GetGeneration(client *storage.Client, ctx context.Context, bucket string, name string) (int64, error) {
	o := client.Bucket(bucket).Object(name)
	objAttrs, err := o.Attrs(ctx)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	gennumber := objAttrs.Generation
	if gennumber > 0 {
		return gennumber, nil
	} else {
		return -1, &MyError{"no generation found"}
	}
}

func FileUpsert(file *bufio.Reader, instance int, fileName string, mime string, user string, uuidString string, objectID int, bucket string, tableName string, endpoint string, adminSecret string, bearer string, createdByInstanceID int, updatedByInstanceID int, constraint string) (int, string, int64, error) {
	UUID := uuid.NewV4().String()

	ustr := ""
	if len(uuidString) > 1 {
		//if there is a uuid header use it else generate one
		ustr = uuidString
		name := strconv.Itoa(instance) + "/" + ustr
		objExists := Objcheck(bucket, name)
		if !objExists {
			ustr = UUID
		} else {
			fmt.Println("found " + name)
		}

	} else {
		ustr = UUID
	}

	objname := strconv.Itoa(instance) + "/" + ustr

	fmt.Println(fileName, objname)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)

		return 0, "", 0, err
	}
	GcsUpload(ctx, client, *file, bucket, objname, mime, fileName)

	objExists := Objcheck(bucket, objname)

	if !objExists {
		fmt.Println("obj doesnt exist")
		return 0, "", 0, err
	}

	gen, err := GetGeneration(client, ctx, bucket, objname)
	if err != nil {
		fmt.Println("failed to get gen")

		return 0, "", 0, err
	} else {
		fmt.Println(gen)
		_ = mime
	}

	id, uuid, err := UpsertAttachmentGen(tableName, fileName, ustr, gen, mime, instance, user, objectID, endpoint, adminSecret, bearer, createdByInstanceID, updatedByInstanceID, constraint)
	if err != nil {
		fmt.Println("could not upsert attachment gen")
		return 0, "", 0, err
	}

	return id, uuid, gen, nil
}

//smartQuery takes the gql query as a string, variables as map string interface and performs hasura query
func SmartQuery(query string, dynamicQuery map[string]interface{}, hasuraEndpoint string, adminSecret string, bearer string) (Responsedata, error) {

	payload := map[string]interface{}{
		"query":     query,
		"variables": dynamicQuery,
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		log.Println("json encoding error", err)
	}

	reqAPI, _ := http.NewRequest("POST", hasuraEndpoint, b)
	reqAPI.Header.Add("content-type", "application/json")

	if len(adminSecret) > 0 {
		reqAPI.Header.Add("X-Hasura-Admin-Secret", adminSecret)
	} else {
		reqAPI.Header.Add("Authorization", bearer)
	}

	resAPI, err := http.DefaultClient.Do(reqAPI)
	rd := Responsedata{}
	if err != nil {
		return rd, err
	}

	bslice, err := ioutil.ReadAll(resAPI.Body)
	bcopy := make([]byte, len(bslice))
	n := copy(bcopy, bslice)
	log.Println("bytes copied", n)
	err = json.Unmarshal(bslice, &rd)
	if err != nil {
		return rd, &MyError{rd.Errors[0].Message}
	}

	if len(rd.Errors) != 0 {
		return rd, &MyError{rd.Errors[0].Message}

	}
	return rd, err
}
