package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clientInstanceError error

type Collection string

const (
	BlogsCollection Collection = "blogs"
)

const (
	Database = "blogs-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {

		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading Env file")
		}

		dbPass := os.Getenv("BACKEND_MONGO_PW")

		clientOptions := options.Client().ApplyURI(dbPass)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client

		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
