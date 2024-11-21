//service/admin_mongoDG.go

package services

import (
	"bookstore/config"
	"bookstore/request"
	"bookstore/response"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	db *mongo.Database
}

func objectIDToInt(oid primitive.ObjectID) int {
	hash := sha1.Sum(oid[:])
	return int(binary.BigEndian.Uint32(hash[:4])) // Use first 4 bytes of the hash
}

// new CreateUser function after using GetNextSequence
func CreateUser(user *request.Request) (*response.Response, error) {
	// Get the next sequence ID for the user collection
	nextID, err := config.GetNextSequence("users")
	if err != nil {
		return nil, fmt.Errorf("failed to generate ID for new user: %v", err)
	}

	// Set the generated ID to the user object
	user.ID = nextID

	// Insert the new user into the "users" collection
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the user document into MongoDB
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// Return the response with the generated ID
	res := &response.Response{
		ID:     user.ID,
		Name:   user.Name,
		Gender: user.Gender,
		Age:    user.Age,
	}

	log.Println("Service close MongoDB")

	return res, nil
}

//old function which was working currectly without giving id
// func CreateUser(movie *request.Request) (*response.Response, error) {
// 	collection := config.DB.Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	result, err := collection.InsertOne(ctx, movie)
// 	if err != nil {
// 		return &response.Response{}, err
// 	}

// 	insertedID := result.InsertedID.(primitive.ObjectID)
// 	intID := objectIDToInt(insertedID) // Call the function here
// 	res := &response.Response{
// 		// ID: insertedID.Hex(),
// 		ID:     intID,
// 		Name:   movie.Name,
// 		Gender: movie.Gender,
// 		Age:    movie.Age,
// 	}
// 	log.Println("Service close MongoDB")

// 	return res, nil
// }

func (d *UserService) GetAllSUsers(recordSize int, offset int, gender string) ([]response.Response, error) {
	log.Println("Service: GetAllSUsers")

	// Set up the MongoDB query filter
	// var filter bson.M
	// if gender != "" {
	// 	filter = bson.M{"gender": gender} // Filter by gender if provided
	// } else {
	// 	filter = bson.M{} // No filter if gender is not provided
	// }

	// Set up the MongoDB query options for pagination
	options := options.Find()
	if recordSize != -1 {
		options.SetLimit(int64(recordSize))
		options.SetSkip(int64(offset))
	}

	// Perform the query on the "users" collection
	collection := config.DB.Collection("users")
	// cursor, err := collection.Find(context.Background(), filter, options)
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Prepare a slice to hold the result documents
	var users []response.Response

	// Iterate through the query results and map them to response objects
	for cursor.Next(context.Background()) {
		var user response.Response
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check if there were any errors while iterating the cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Return the result
	return users, nil
}

// GetUserByID retrieves a user by their ID from the MongoDB users collection.
func (d *UserService) GetUserByID(id string) (*response.Response, error) {
	log.Println("Service: GetUserByID")

	// Convert string ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Set up MongoDB query filter
	filter := bson.M{"_id": objectID}

	// Perform the query on the "users" collection
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user response.Response

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	// Return the user document
	return &user, nil
}
