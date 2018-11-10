package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var projectID = os.Getenv("PROJECT_ID")

func main() {
	ctx := context.Background()

	bucket, err := createBucket(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := bucket.upload(ctx, "./aqua.png", "image"); err != nil {
		log.Fatal(err)
		return
	}

	if err := bucket.download(ctx, "image", "./aqua_downloaded.png"); err != nil {
		log.Fatal(err)
		return
	}
}

func getClient(ctx context.Context) (*storage.Client, error) {
	opt := option.WithCredentialsFile(os.Getenv("GCLOUD_CRENTIAL_FILE_PATH"))
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %v", err)
	}
	return client, nil
}

// GCSBucket は Google Cloud Storage のバケットを表すstruct
type GCSBucket struct {
	handle *storage.BucketHandle
}

func createBucket(ctx context.Context) (*GCSBucket, error) {
	client, err := getClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get client: %v", err)
	}

	bucket := client.Bucket("go-client-sandbox")
	attr, err := bucket.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get bucket attributes: %v", err)
	}
	// Create Bucket if not exist
	if attr == nil {
		if err := bucket.Create(ctx, projectID, nil); err != nil {
			return nil, fmt.Errorf("Failed to create bucket: %v", err)
		}
	}

	return &GCSBucket{bucket}, nil
}

func (b *GCSBucket) upload(ctx context.Context, filepath, name string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v", err)
	}

	object := b.handle.Object(name)
	objectWriter := object.NewWriter(ctx)
	if _, err := objectWriter.Write(file); err != nil {
		return fmt.Errorf("Failed to upload: %v", err)
	}
	defer objectWriter.Close()

	return nil
}

func (b *GCSBucket) download(ctx context.Context, name, filepath string) error {
	object := b.handle.Object(name)
	objectReader, err := object.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Failed to download: %v", err)
	}
	defer objectReader.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %v", err)
	}
	defer file.Close()

	r := bufio.NewReader(objectReader)
	_, err = r.WriteTo(file)
	if err != nil {
		return fmt.Errorf("Failed to write file: %v", err)
	}

	return nil
}
