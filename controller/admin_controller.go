// controller/admin_controller.go

package controller

import (
	"database/sql"
	// "bookstore/response"

	// "bookstore/manager"
	"bookstore/manager"
	"bookstore/request"

	"log"
	"net/http"
	"strconv"

	// "bookstore/model"
	"github.com/labstack/echo/v4"
)

// Declare the AdminManager
// var movieManager *services.AdminService

//InitializeController initializes the controller with a AdminService
// func InitializeController(service *services.AdminService) {
// 	movieManager = service
// }

var movieManager *manager.AdminManager
var userManager *manager.UserManager

func SetManagers(uManager *manager.UserManager) {
	userManager = uManager
}

// var movieService *services.AdminService

func InitializeController(mgr *manager.AdminManager) {
	movieManager = mgr
}

// func InitializeControllerService(service *services.AdminService) {
// 	movieService = service
// }

// Create the Admin table
// err := model.CreateAdminTable(db);
// if  err != nil {
// 	log.Fatalf("Failed to create Admin table: %v", err)
// }

// CreateAdmin handles the creation of a new movie
func CreateAdmin(c echo.Context) error {
	log.Println("ctr begin")

	req := new(request.Request)

	// Bind the request data to the Admin struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Validate the movie data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	flag := c.QueryParam("flag")

	// Call the AdminManager to create the movie
	r, err := movieManager.CreateAdmin(req, flag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Println("ctr close")

	// Return the response directly from the service/manager
	return c.JSON(http.StatusCreated, r)
}

func GetAllAdmins(c echo.Context) error {
	//Converting QueryParams String to Interger
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1 // Default to page 1 if invalid
	}

	recordSize, err := strconv.Atoi(c.QueryParam("recordsize"))
	if err != nil || recordSize <= 0 {
		recordSize = recordSize * 1
	}

	if page == -1 {
		recordSize = -1 // Indicating to fetch all records
	}
	gender := c.QueryParam("gender")
	flag := c.QueryParam("flag")

	// Create a UserManager instance (ensure this is initialized properly)
	// userManager := manager.NewUserManager(userService) // userService must be initialized elsewhere

	// Call the Manager to fetch paginated movies
	movies, err := movieManager.GetAllAdmins(userManager, page, recordSize, gender, flag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, movies)
}

// GetAdmin fetches a movie by ID
func GetAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	movie, err := movieManager.GetAdminByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Admin not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// res := response.Response{
	// 	ID:    movie.ID,
	// 	Name:  movie.Name,
	// 	Genre: movie.Genre,
	// 	Price: movie.Price,
	// }
	return c.JSON(http.StatusOK, movie)
}

// UpdateAdmin handles updating an existing movie
func UpdateAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	req := new(request.Request)

	// Bind the request data to the Admin struct
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	req.ID = id

	// Validate the movie data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call the AdminManager to update the movie
	err = movieManager.UpdateAdmin(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, req)
}

// DeleteAdmin deletes a movie by ID
func DeleteAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
	}

	err = movieManager.DeleteAdmin(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Admin deleted successfully"})
}
