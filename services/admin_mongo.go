//service/admin_mongoDG.go

package services

import (
	"bookstore/config"
	"bookstore/request"
	"bookstore/response"

	// "bookstore/services"
	"context"

	// "crypto/sha1"
	// "encoding/binary"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	db *mongo.Database
}

// new CreateUser function after using GetNextSequence
func CreateUser(user *request.Request) (*response.Response, error) {
	// Get the next sequence ID for the "users" collection
	nextID, err := IncrementMongoId() // Now passing "users"
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
	_, err = collection.InsertOne(ctx, user) // Don't need to capture result if not using it
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

func (d *UserService) GetAllSUsers(recordSize int, offset int, gender string) ([]response.Response, error) {
	log.Println("Service: GetAllSUsers")

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
func (d *UserService) GetUserByID(id int) (*response.Response, error) {
	log.Println("Service: GetUserByID")

	// Set up MongoDB query filter
	filter := bson.M{"_id": id}

	// Perform the query on the "users" collection
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user response.Response

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	// Return the user document
	return &user, nil
}
func (d *UserService) UpdateUser(req *request.Request) error {
	log.Println("Service: UpdateUser")

	// Set up MongoDB query filter
	filter := bson.M{"_id": req.ID}

	// Specify the update operation
	update := bson.M{"$set": bson.M{
		"ID":     req.ID,
		"Name":   req.Name,
		"Gender": req.Gender,
		"age":    req.Age,
	}}

	// Perform the update operation on the "users" collection
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	// Check if any document was updated
	if result.MatchedCount == 0 {
		return fmt.Errorf("no user found with ID %d", req.ID)
	}

	log.Printf("User with ID %d successfully updated", req.ID)
	return nil
}
func (d *UserService) DeleteUser(id int) error {
	log.Println("Service: DeleteUser")

	// Set up MongoDB query filter
	filter := bson.M{"_id": id}

	// Perform the delete operation on the "users" collection
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("no user found with ID %s", id)
	}

	log.Printf("User with ID %s successfully deleted", id)
	return nil
}

func IncrementMongoId() (int, error) {
	collection := config.DB.Collection("counters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.M{"_id": "restaurant_id"}, bson.M{"$inc": bson.M{"seq": 1}})
	if err != nil {
		return 0, fmt.Errorf("there is an error in id updation:%v", err)
	}
	if res.ModifiedCount == 0 {
		return 0, fmt.Errorf("there is an error in increament mongodb ID")
	}
	var result response.IdResponse
	err = collection.FindOne(ctx, bson.M{"_id": "restaurant_id"}).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("error in getting sequence: %v", err)
	}
	return result.ID, nil
}

func DecrementMongoId() error {
	collection := config.DB.Collection("menu")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.M{"id": "restaurant_id"}, bson.M{"seq": bson.M{"$inc": -1}})
	if err != nil {
		return fmt.Errorf("there is an error in updation:%v", err)
	}
	if res.ModifiedCount == 0 {
		return fmt.Errorf("there is an error in increament mongodb ID")
	}
	return nil
}
