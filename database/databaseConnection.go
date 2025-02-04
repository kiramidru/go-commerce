package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	url := os.Getenv("MONGODB_URL")
	if url == "" {
		log.Fatal("No environment found")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, error := mongo.Connect(context.TODO(), opts)

	if error != nil {
		log.Fatal(error)
	}

	defer func() {
		if error = client.Disconnect(context.TODO()); error != nil {
			log.Fatal(error)
		}
	}()
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("e-commerce").Collection(collectionName)
	return collection
}
