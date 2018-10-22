package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Message struct {
	Name          string `json:"name"`
	ProfilePicUrl string `json:profilePicUrl`
	Text          string `json:"text"`
}

func main() {
	ctx := context.Background()

	opt := option.WithCredentialsFile(os.Getenv("GCLOUD_CRENTIAL_FILE_PATH"))
	config := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DATASTORE_URL"),
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Get
	var messages map[string]Message
	err = client.NewRef("messages").Get(ctx, &messages)
	if err != nil {
		log.Fatal(err)
	}

	for key, message := range messages {
		log.Printf("%s: %+v\n", key, message)
	}

	// Push
	newMessage := Message{
		Name: "golang client",
		Text: "this is message by golang client",
	}
	_, err = client.NewRef("messages").Push(ctx, newMessage)
	if err != nil {
		log.Fatal(err)
	}
}
