package model

import (
	"database/sql"
	"fmt"
)

// import (
// 	"github.com/go-playground/validator/v10"
// )

// var validate = validator.New()

type Admin struct {
	ID     int     `json:"id"`
	Name   string  `json:"name" validate:"required,min=1,max=100"`
	Gender string  `json:"gender" validate:"required,min=1,max=50"`
	Age    float64 `json:"age" validate:"required,gt=0"`
}

func CreateAdminTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS Admin (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		gender VARCHAR(50) NOT NULL,
		age NUMERIC NOT NULL CHECK (age > 0)
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating Admin table: %w", err)
	}

	return nil
}

// // Validate the Admin struct
// func (m *Admin) Validate() error {
// 	return validate.Struct(m)
// }

// // Create a new admin in the database
// func CreateAdmin(db *sql.DB, admin *Admin) error {
// 	if err := admin.Validate(); err != nil {
// 		return err
// 	}
// 	return db.QueryRow("INSERT INTO postgres (name, gender, age) VALUES ($1, $2, $3) RETURNING id",
// 		admin.Name, admin.Gender, admin.Age).Scan(&admin.ID)
// }

// // Get a admin by ID
// func GetAdminByID(db *sql.DB, id int) (*Admin, error) {
// 	admin := &Admin{}
// 	err := db.QueryRow("SELECT id, name, gender, age FROM book_admins WHERE id = $1", id).
// 		Scan(&admin.ID, &admin.Name, &admin.Gender, &admin.Age)
// 	if err == sql.ErrNoRows {
// 		return nil, errors.New("admin not found")
// 	}
// 	return admin, err
// }

// // Get all book_admins from the database
// func GetAllAdmins(db *sql.DB) ([]Admin, error) {
// 	rows, err := db.Query("SELECT id, name, gender, age FROM book_admins")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var book_admins []Admin
// 	for rows.Next() {
// 		var admin Admin
// 		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Gender, &admin.Age); err != nil {
// 			return nil, err
// 		}
// 		book_admins = append(book_admins, admin)
// 	}
// 	return book_admins, nil
// }

// // Update a admin's details
// func UpdateAdmin(db *sql.DB, admin *Admin) error {
// 	if err := admin.Validate(); err != nil {
// 		return err
// 	}
// 	_, err := db.Exec("UPDATE book_admins SET name=$1, gender=$2, age=$3 WHERE id=$4",
// 		admin.Name, admin.Gender, admin.Age, admin.ID)
// 	return err
// }

// // Delete a admin by ID
// func DeleteAdmin(db *sql.DB, id int) error {
// 	_, err := db.Exec("DELETE FROM book_admins WHERE id=$1", id)
// 	return err
// }

// package model

// import (
// 	"database/sql"
// 	"errors"
// 	// "github.com/go-playground/validator/v10"
// )

// type Admin struct {
// 	ID    int     `json:"id"`
// 	Name  string  `json:"name"`
// 	Gender string  `json:"gender"`
// 	Age float64 `json:"age"`
// }

// // Create a new admin in the database
// func CreateAdmin(db *sql.DB, admin *Admin) error {
// 	return db.QueryRow("INSERT INTO book_admins (name, gender, age) VALUES ($1, $2, $3) RETURNING id",
// 		admin.Name, admin.Gender, admin.Age).Scan(&admin.ID)
// }

// // Get a admin by ID
// func GetAdminByID(db *sql.DB, id int) (*Admin, error) {
// 	admin := &Admin{}
// 	err := db.QueryRow("SELECT id, name, gender, age FROM book_admins WHERE id = $1", id).
// 		Scan(&admin.ID, &admin.Name, &admin.Gender, &admin.Age)
// 	if err == sql.ErrNoRows {
// 		return nil, errors.New("admin not found")
// 	}
// 	return admin, err
// }

// // Get all book_admins from the database
// func GetAllAdmins(db *sql.DB) ([]Admin, error) {
// 	rows, err := db.Query("SELECT id, name, gender, age FROM book_admins")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var book_admins []Admin
// 	for rows.Next() {
// 		var admin Admin
// 		if err := rows.Scan(&admin.ID, &admin.Name, &admin.Gender, &admin.Age); err != nil {
// 			return nil, err
// 		}
// 		book_admins = append(book_admins, admin)
// 	}
// 	return book_admins, nil
// }

// // Update a admin's details
// func UpdateAdmin(db *sql.DB, admin *Admin) error {
// 	_, err := db.Exec("UPDATE book_admins SET name=$1, gender=$2, age=$3 WHERE id=$4",
// 		admin.Name, admin.Gender, admin.Age, admin.ID)
// 	return err
// }

// // Delete a admin by ID
// func DeleteAdmin(db *sql.DB, id int) error {
// 	_, err := db.Exec("DELETE FROM book_admins WHERE id=$1", id)
// 	return err
// }
