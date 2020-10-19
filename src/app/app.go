package app

import (
	"context"
    "fmt"
	"log"
	"net/http"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Client : This is the client for Mongo
var Client = SetupDB()

// GetClient : Returns the Mongo client
func GetClient() *mongo.Client {
	return Client
}

// Run : Initializer for the API
func Run(PORT string) {
	log.Printf("Listening for requests at http://localhost:%v", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}

// SetupDB : Sets up the database connection
func SetupDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}