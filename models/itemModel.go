package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson: "_product_id"`
	Name        string             `json: "name"`
	Description string             `json: "description"`
	Quantity    string             `json: "quantity"`
	Product_id  string             `json: "item_id"`
}
