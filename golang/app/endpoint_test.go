package app

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


// Setup a collection for the test
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

// Truncate the collection
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

// Fixture
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

// Test setup for the /documents route
var documentTests = []struct {
	in  string
	out []doc
}{
	{"/documents", fixture},
	{"/documents/Bar", fixture[1:3]},
	{"/documents/bar", fixture[1:3]},
	{"/documents/Luise", fixture[2:3]},
	{"/documents/john", fixture[0:2]},
}

func TestDocumentsRoute(t *testing.T) {

	for _, tt := range documentTests {
		// Setup fixture
		router := SetupRouter()
		setupDb(fixture)

		// Test
		t.Run(tt.in, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.in, nil)
			router.ServeHTTP(w, req)

			// Assertions
			var actResult docResponse
			var expResult = docResponse {
				Documents: tt.out,
			}
			json.Unmarshal([]byte(w.Body.String()), &actResult)

			if ! cmp.Equal(actResult, expResult) {
				fmt.Println("actResult", actResult)
				fmt.Println("expResult", expResult)
				t.Errorf("Result not equal to expected results")
			}

			assert.Equal(t, 200, w.Code)
		})

		// Cleanup
		cleanupDb()
	}
}
