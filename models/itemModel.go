package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID `bson: "id"`
	Name        string             `json: "name"`
	Description string             `json: "description"`
	Quantity    string             `json: "quantity"`
}
