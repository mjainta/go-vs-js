package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type doc struct {
	ID 		primitive.ObjectID `bson:"_id"`
	Name 	string						 `bson:"name"`
}

func connect() (*mongo.Client, context.Context) {

	// Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb://admin:foobar@localhost"
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

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

func getAllDocuments() []*doc {
	var documents []*doc
	var client, ctx = connect()
	defer disconnect(client, ctx)

	// Get the collection documents iterate through them
	collection := client.Database("mydatabase").Collection("mycoll")
	cur, err := collection.Find(ctx, bson.D{})

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

		fmt.Println(elem)
		documents = append(documents, &elem)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", documents)
	return documents
}