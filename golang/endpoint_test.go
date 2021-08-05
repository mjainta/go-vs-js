package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDb(fixture []doc) {
	var dbUri = os.Getenv("DB_URI")
	var dbName = os.Getenv("DB_NAME")
	var collName = os.Getenv("DB_COLL")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))

	if err != nil {
		panic(err)
	}

	collection := client.Database(dbName).Collection(collName)
	for _, entry := range fixture {
		_, err := collection.InsertOne(ctx, entry)

		if err != nil {
			panic(err)
		}
	}
}

func cleanupDb() {
	var dbUri = os.Getenv("DB_URI")
	var dbName = os.Getenv("DB_NAME")
	var collName = os.Getenv("DB_COLL")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))

	if err != nil {
		panic(err)
	}

	collection := client.Database(dbName).Collection(collName)
	_, err = collection.DeleteMany(ctx, bson.M{})

	if err != nil {
		panic(err)
	}
}

func TestDocumentsRoute(t *testing.T) {
	router := setupRouter()
	var fixture = []doc {
		{
			ID: primitive.NewObjectID(),
			Name: "John Doe",
		},
		{
			ID: primitive.NewObjectID(),
			Name: "John Bar",
		},
		{
			ID: primitive.NewObjectID(),
			Name: "Luise Bar",
		},
	}
	setupDb(fixture)
	defer cleanupDb()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents", nil)
	router.ServeHTTP(w, req)

	var actResult docResponse
	var expResult = docResponse {
		Documents: fixture,
	}
	json.Unmarshal([]byte(w.Body.String()), &actResult)

	if ! cmp.Equal(actResult, expResult) {
		fmt.Println("actResult", actResult)
		fmt.Println("expResult", expResult)
		t.Errorf("Result not equal to expected results")
	}

	assert.Equal(t, 200, w.Code)
}

func TestDocumentsFilteredRoute(t *testing.T) {
	router := setupRouter()
	var fixture = []doc {
		{
			ID: primitive.NewObjectID(),
			Name: "John Doe",
		},
		{
			ID: primitive.NewObjectID(),
			Name: "John Bar",
		},
		{
			ID: primitive.NewObjectID(),
			Name: "Luise Bar",
		},
	}
	setupDb(fixture)
	defer cleanupDb()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents/Bar", nil)
	router.ServeHTTP(w, req)

	var actResult docResponse
	var expResult = docResponse {
		Documents: fixture[1:3],
	}
	json.Unmarshal([]byte(w.Body.String()), &actResult)

	if ! cmp.Equal(actResult, expResult) {
		fmt.Println("actResult", actResult)
		fmt.Println("expResult", expResult)
		t.Errorf("Result not equal to expected results")
	}

	assert.Equal(t, 200, w.Code)
}
