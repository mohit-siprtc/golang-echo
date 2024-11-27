// config.go
package db_conn

import (
	"bookstore/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Load environment variables from .env file
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Connect to MongoDB and return the client
func ConnectDB() {
	loadEnv() // Load environment variables from .env

	mongoConfig := config.LoadMongoConfig()

	mongoURI := mongoConfig.MONGODB_URI
	mongo_DB_NAME := mongoConfig.MONGO_DB_NAME

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
	DB = client.Database(mongo_DB_NAME)
	fmt.Println("Connected to MongoDB")
	CreateCounterSeq()
	// return client
	// DB = client.Database(os.Getenv("MONGO_DB_NAME"))
	fmt.Println(DB)
}

// GetCollection retrieves a collection by name from MongoDB
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	mongoConfig := config.LoadMongoConfig()
	dbName := mongoConfig.MONGO_DB_NAME
	if dbName == "" {
		log.Fatal("MONGO_DB_NAME not set in .env")
	}
	return client.Database(dbName).Collection(collectionName)
}
func CreateCounterSeq() error {
	collection := DB.Collection("counters") // Counter collection name

	// Check if the sequence document exists
	var result bson.M
	err := collection.FindOne(context.Background(), bson.M{"_id": "restaurant_id"}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Sequence document doesn't exist, create it with an initial value of 0
		_, err = collection.InsertOne(context.Background(), bson.M{
			"_id": "restaurant_id", // Unique identifier for the sequence
			"seq": 0,               // Initial sequence value
		})
		if err != nil {
			return fmt.Errorf("error creating counter sequence document: %v", err)
		}
		fmt.Println("Counter sequence document created with initial value.")
	} else if err != nil {
		return fmt.Errorf("error checking counter sequence document: %v", err)
	} else {
		fmt.Println("Counter sequence document already exists.")
	}
	return nil
}
