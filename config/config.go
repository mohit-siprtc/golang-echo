// config.go
package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// GetNextSequence function retrieves the next sequence number
func GetNextSequence() (int, error) {
	collection := DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get the last inserted user document (assuming ID is an integer)
	var lastUser struct {
		ID int `bson:"id"`
	}

	// Find the last inserted user document sorted by ID in descending order
	err := collection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"id": -1})).Decode(&lastUser)
	if err != nil && err != mongo.ErrNoDocuments {
		return 0, fmt.Errorf("failed to retrieve last user: %v", err)
	}

	// If no user exists, start the ID from 1
	if err == mongo.ErrNoDocuments {
		return 1, nil
	}

	// Return the next sequence (last ID + 1)
	return lastUser.ID + 1, nil
}

// Load environment variables from .env file
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Connect to MongoDB and return the client
func ConnectDB() {
	loadEnv() // Load environment variables from .env

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database(os.Getenv("MONGO_DB_NAME"))
	fmt.Println("Connected to MongoDB")
	// return client
	// DB = client.Database(os.Getenv("MONGO_DB_NAME"))
	fmt.Println(DB)
}

// GetCollection retrieves a collection by name from MongoDB
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		log.Fatal("MONGO_DB_NAME not set in .env")
	}
	return client.Database(dbName).Collection(collectionName)
}
