package rdgo

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"cloud.google.com/go/storage"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

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
		fmt.Println(rd.Data.Attachment)
		if len(rd.Data.ContractAttachment.Returning) > 0 {
			return rd.Data.ContractAttachment.Returning[0].Id, rd.Data.ContractAttachment.Returning[0].UUID, nil
		} else if len(rd.Data.PoHeaderAttachment.Returning) > 0 {
			return rd.Data.PoHeaderAttachment.Returning[0].Id, rd.Data.PoHeaderAttachment.Returning[0].UUID, nil
		}
	}
	return -1, "", err
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
			fmt.Println("found " + UUID)
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
