package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Name string `firestore:"name,omitempty"`
}

func TestDocumentCRUD(t *testing.T) {
	ctx := context.Background()

	client, err := createClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Set
	set := &User{Name: "Tom"}
	_, err = client.Collection("users").Doc("1").Set(ctx, set)
	assert.NoError(t, err)

	// Get
	d, err := client.Doc("users/1").Get(ctx)
	assert.NoError(t, err)
	assert.True(t, d.Exists())

	actual := &User{}
	d.DataTo(&actual)
	require.Equal(t, set.Name, actual.Name)

	// Update
	newName := "Bob"
	set.Name = newName
	_, err = client.Doc("users/1").Set(ctx, set)
	assert.NoError(t, err)

	d, err = client.Doc("users/1").Get(ctx)
	assert.NoError(t, err)
	assert.True(t, d.Exists())

	actual = &User{}
	d.DataTo(&actual)
	require.Equal(t, newName, actual.Name)

	// Delete
	_, err = client.Doc("users/1").Delete(ctx)
	assert.NoError(t, err)

	d, err = client.Doc("users/1").Get(ctx)
	require.Error(t, err) // firestore.errNilDocRef が発生する
	// https://godoc.org/cloud.google.com/go/firestore#DocumentRef.Get
	// > Get retrieves the document. If the document does not exist, Get return a NotFound error, which can be checked with
	// > grpc.Code(err) == codes.NotFound
	// grpc.Code メソッドは deprecated なので、 status.Code メソッドを利用
	require.Equal(t, status.Code(err), codes.NotFound)
	require.False(t, d.Exists())
}
