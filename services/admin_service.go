// services/admin_service.go

package services

import (
	"bookstore/request"
	"bookstore/response"
	"database/sql"
	"log"
)

// var res  *response.Response

// func res_ctrl(r *response.Response) {
// 	res = r
// }

// AdminService provides database-related operations for movies
type AdminService struct {
	DB *sql.DB
}

// NewAdminService initializes a new AdminService
func NewAdminService(db *sql.DB) *AdminService {
	return &AdminService{DB: db}
}

// CreateAdmin adds a new movie to the database
func (s *AdminService) CreateAdmin(movie *request.Request) (*response.Response, error) {
	log.Println("Service begin")

	var id int
	err := s.DB.QueryRow("INSERT INTO Admin (name, gender, age) VALUES ($1, $2, $3) RETURNING id", movie.Name, movie.Gender, movie.Age).Scan(&id)
	log.Println("err-------------->", err)
	if err != nil {
		return nil, err
	}

	res := &response.Response{
		ID:     id,
		Name:   movie.Name,
		Gender: movie.Gender,
		Age:    movie.Age,
	}

	log.Println("Service close")
	return res, nil
}

// services/movie_service.go

// services/movie_service.go

func (s *AdminService) GetAllAdmins(recordSize, offset int, gender string) ([]response.Response, error) {
	log.Println("Service: GetAllAdmins")

	var rows *sql.Rows
	var err error

	// Determine the query based on the genre filter and pagination settings
	if recordSize == -1 {
		// Fetch all records
		if gender != "" {
			log.Println("Querying with gender filter (fetch all):", gender)
			rows, err = s.DB.Query("SELECT id, name, gender, price FROM Admin WHERE genre = $1", gender)

		} else {
			rows, err = s.DB.Query("SELECT id, name, gender, age FROM Admin")
			// return rows

		}

	} else {
		// Fetch with pagination (LIMIT and OFFSET)
		if gender != "" {
			log.Println("Querying with gender filter and pagination:", gender, recordSize, offset)
			rows, err = s.DB.Query("SELECT id, name, gender, age FROM Admin WHERE gender = $1 LIMIT $2 OFFSET $3", gender, recordSize, offset)

		} else {
			rows, err = s.DB.Query("SELECT id, name, gender, age FROM Admin LIMIT $1 OFFSET $2", recordSize, offset)
		}
	}

	// Check if there was an error executing the query
	if err != nil {
		return nil, err
	}
	// defer rows.Close()

	// Process query results
	var movies []response.Response
	for rows.Next() {
		var movie response.Response
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Gender, &movie.Age); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	// Check for any error encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Println("------------Postgres GetAll Working-----------")
	return movies, nil
}

// func (s *AdminService) GetAllAdmins(recordSize, offset int,genre string) ([]response.Response, error) {
// 	log.Println("Service: GetAllAdmins")

// 	var rows *sql.Rows
// 	// var err error

// 	// If recordSize is -1, fetch all records
// 	if recordSize == -1 {
// 		rows, err := s.DB.Query("SELECT id, name, genre, price FROM movies")
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer rows.Close()

// 		var movies []response.Response
// 		for rows.Next() {
// 			var movie response.Response
// 			if err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price); err != nil {
// 				return nil, err
// 			}
// 			movies = append(movies, movie)
// 		}
// 		return movies, nil
// 	} else {
// 		// If genre filter is provided, use WHERE clause to filter by genre with pagination
// 		if genre != "" {
// 			rows, err := s.DB.Query("SELECT id, name, genre, price FROM movies WHERE genre = $1 LIMIT $2 OFFSET $3", genre, recordSize, offset)
// 			if err != nil {
// 				return nil, err
// 			}
// 			defer rows.Close()
// 		} else {
// 			rows, err := s.DB.Query("SELECT id, name, genre, price FROM movies LIMIT $1 OFFSET $2", recordSize, offset)
// 			if err != nil {
// 				return nil, err
// 			}
// 			defer rows.Close()
// 		}
// 	}

// 	// defer rows.Close()
// 	// Query the database with pagination (LIMIT and OFFSET)
// 	// rows, err = s.DB.Query("SELECT id, name, genre, price FROM movies LIMIT $1 OFFSET $2", recordSize, offset)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// defer rows.Close()

//		var movies []response.Response
//		for rows.Next() {
//			var movie response.Response
//			if err := rows.Scan(&movie.ID, &movie.Name, &movie.Genre, &movie.Price); err != nil {
//				return nil, err
//			}
//			movies = append(movies, movie)
//		}
//		return movies, nil
//	}
//
// GetAdminByID retrieves a single movie by ID
func (s *AdminService) GetAdminByID(id int) (*response.Response, error) {
	var movie response.Response
	err := s.DB.QueryRow("SELECT id, name, gender, age FROM Admin WHERE id = $1", id).
		Scan(&movie.ID, &movie.Name, &movie.Gender, &movie.Age)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// UpdateAdmin modifies an existing movie
func (s *AdminService) UpdateAdmin(movie *request.Request) error {
	_, err := s.DB.Exec("UPDATE Admin SET name=$1, gender=$2, age=$3 WHERE id=$4",
		movie.Name, movie.Gender, movie.Age, movie.ID)
	return err
}

// DeleteAdmin removes a movie by ID
func (s *AdminService) DeleteAdmin(id int) error {
	_, err := s.DB.Exec("DELETE FROM Admin WHERE id=$1", id)
	return err
}
