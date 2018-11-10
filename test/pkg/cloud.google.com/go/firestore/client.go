package main

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func createClient(ctx context.Context) (*firestore.Client, error) {
	projectID := os.Getenv("PROJECT_ID")
	opt := option.WithCredentialsFile(os.Getenv("GCLOUD_CRENTIAL_FILE_PATH"))

	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		return nil, err
	}
	return client, err
}
