package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"bookstore/controller"
	"bookstore/manager"
	"bookstore/model"
	"bookstore/route"
	"bookstore/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	// "go.mongodb.org/mongo-driver/mongo"

	"bookstore/config"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Database connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := model.CreateAdminTable(db); err != nil {
		log.Fatalf("Failed to create Admin table: %v", err)
	}

	fmt.Println("Admin table created successfully! postgress")

	// Initialize Service, Manager, and Controller

	movieService := services.NewAdminService(db)          // Service handles database operations
	adminManager := manager.NewAdminManager(movieService) // Manager handles business logic
	controller.InitializeController(adminManager)         // Controller handles HTTP requests

	config.ConnectDB()

	userService := &services.UserService{}
	userManager := manager.NewUserManager(userService)

	controller.SetManagers(userManager)
	// MongoDB connection setup
	// mongoClient := config.ConnectDB()                             // Assuming you have a ConnectDB() function in your config package
	// usersCollection := config.GetCollection(mongoClient, "users") // Get "users" collection from MongoDB
	// booksCollection := config.GetCollection(mongoClient, "books") // Get "books" collection from MongoDB
	// fmt.Println(usersCollection)

	// Create Echo instance and setup routes
	e := echo.New()

	route.SetupRoutes(e) // Setup routes using the route package

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
