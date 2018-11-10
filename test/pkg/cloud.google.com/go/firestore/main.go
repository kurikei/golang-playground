package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/iterator"
)

// Message はチャットのメッセージを表すstruct
type Message struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func main() {
	ctx := context.Background()

	client, err := createClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// メッセージを追加する
	newMessage := Message{
		Name: "Firestore Golang Library",
		Text: fmt.Sprintf("This is message by %v", time.Now().Unix()),
	}
	_, _, err = client.Collection("messages").Add(ctx, &newMessage)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	// メッセージを取得する
	iter := client.Collection("messages").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
