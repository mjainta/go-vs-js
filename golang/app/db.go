package app

import (
	"os"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func connect() (*mongo.Client, context.Context) {

	var dbUri = os.Getenv("DB_URI")

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))

	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	return client, ctx
}

func disconnect(client *mongo.Client, ctx context.Context) {
	client.Disconnect(ctx)
}

func getDocuments(filter docFilter) []*doc {
	var queryFilter = bson.M{}

	if filter.Type == "name" {
		queryFilter = bson.M{
			"name": bson.M{
				"$regex": primitive.Regex{
					Pattern: fmt.Sprintf("%s", filter.Value),
					Options: "i",
				},
			},
		}
	}

	var documents []*doc
	var client, ctx = connect()
	defer disconnect(client, ctx)

	var dbName = os.Getenv("DB_NAME")
	var collName = os.Getenv("DB_COLL")
	// Get the collection documents iterate through them
	collection := client.Database(dbName).Collection(collName)
	cur, err := collection.Find(ctx, queryFilter)

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var elem doc
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}

		documents = append(documents, &elem)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", documents)
	return documents
}
