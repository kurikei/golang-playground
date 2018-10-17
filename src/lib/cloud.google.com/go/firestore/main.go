package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func main() {
	ctx := context.Background()

	projectID := os.Getenv("PROJECT_ID")
	opt := option.WithCredentialsFile(os.Getenv("GCLOUD_CRENTIAL_FILE_PATH"))

	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	newMessage := Message{
		Name: "Firestore Golang Library",
		Text: fmt.Sprintf("This is message by %v", time.Now().Unix()),
	}
	_, _, err = client.Collection("messages").Add(ctx, &newMessage)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}
