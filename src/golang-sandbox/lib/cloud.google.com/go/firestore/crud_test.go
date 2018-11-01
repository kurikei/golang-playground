package main

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Name string `firestore:"name,omitempty"`
}

type Type struct {
	String string                 `firestore:"string,omitempty"`
	Bool   bool                   `firestore:"bool,omitempty"`
	Float  float64                `firestore:"float,omitempty"`
	Time   time.Time              `firestore:"time,omitempty"`
	Array  []interface{}          `firestore:"array,omitempty"`
	Null   interface{}            `firestore:"null,omitempty"`
	Map    map[string]interface{} `firestore:"map,omitempty"`
	DocRef *firestore.DocumentRef `firestore:"doc_ref,omitempty"`
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

func TestCollectionNewDoc(t *testing.T) {
	ctx := context.Background()

	client, err := createClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	dRef := client.Collection("users").NewDoc()
	u := &User{Name: "Alice"}
	_, err = dRef.Set(ctx, u)
	assert.NoError(t, err)

	d, err := client.Collection("users").Doc(dRef.ID).Get(ctx)
	assert.NoError(t, err)
	actual := &User{}
	d.DataTo(actual)
	require.Equal(t, u.Name, actual.Name)
}

func TestType(t *testing.T) {
	ctx := context.Background()

	client, err := createClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// 参照型を確かめるために user ドキュメントのを作成する
	u := &User{Name: "Tom"}
	uDocRef, _, err := client.Collection("users").Add(ctx, u)
	assert.NoError(t, err)

	typeObj := &Type{
		String: "hoge",
		Bool:   true,
		Float:  3.14,
		Time:   time.Now(),
		Array:  []interface{}{1, "string", true},
		Null:   nil,
		Map:    map[string]interface{}{"key": "value"},
		DocRef: uDocRef,
	}
	typeRef, _, err := client.Collection("types").Add(ctx, typeObj)
	assert.NoError(t, err)

	// firestoreから取得したデータの参照から実態が引いてこれるか確認
	ts, err := typeRef.Get(ctx)
	assert.NoError(t, err)
	newType := &Type{}
	err = ts.DataTo(newType)
	assert.NoError(t, err)

	actualUser := &User{}
	us, err := newType.DocRef.Get(ctx)
	us.DataTo(actualUser)
	require.Equal(t, u.Name, actualUser.Name)
}
