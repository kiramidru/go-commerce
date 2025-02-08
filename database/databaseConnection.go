package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	url := os.Getenv("MONGODB_URL")

	if url == "" {
		log.Fatal("No environment file found")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("e-commerce").Collection(collectionName)
	return collection
}
