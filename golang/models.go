package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type doc struct {
	ID 		primitive.ObjectID `bson:"_id"`
	Name 	string						 `bson:"name"`
}

type docResponse struct {
	Documents []doc `json:"documents"`
}
