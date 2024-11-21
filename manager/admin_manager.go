// manager/admin_manager.go

package manager

import (
	"bookstore/request"
	"bookstore/response"
	"bookstore/services"
	"fmt"
	"strconv"

	// "fmt"
	"log"
)

// AdminManager handles business logic for Admins
type AdminManager struct {
	AdminService *services.AdminService
}

// NewAdminManager initializes a new AdminManager
func NewAdminManager(AdminService *services.AdminService) *AdminManager {
	return &AdminManager{AdminService: AdminService}
}

type UserManager struct {
	UserService *services.UserService
}

func NewUserManager(UserService *services.UserService) *UserManager {
	return &UserManager{UserService: UserService}
}

// CreateAdmin adds a new Admin by interacting with the service
func (m *AdminManager) CreateAdmin(Admin *request.Request, flag string) (*response.Response, error) {
	log.Println("Mgr begin")
	log.Println(Admin)
	log.Println("flag-------------->", flag)
	if flag == "true" {
		res, err := m.AdminService.CreateAdmin(Admin)
		if err != nil {
			return nil, err
		}
		return res, nil

	} else {
		res, err := services.CreateUser(Admin)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	// fmt.Println("Mgr close")
	// return &response.Response{}, nil

}

// GetAllAdmins retrieves all Admins by calling the service
// func (m *AdminManager) GetAllAdmins() ([]response.Response, error) {
// 	return m.AdminService.GetAllAdmins()
// }

// func (m *AdminManager) GetAllAdmins(limit, offset int) ([]response.Response, error) {
// 	// Here you can implement any additional business logic
// 	log.Println("Manager: GetAllAdmins")
// 	return m.AdminService.GetAllAdmins(limit, offset)
// }

// func (m *AdminManager) GetAllAdmins(page, recordSize int, gender string, flag string) ([]response.Response, error) {

// 	if flag == "true" {
// 		// If page is -1, set recordSize to -1 to fetch all records
// 		if page == -1 {
// 			recordSize = -1
// 		}

// 		// Calculate offset if not fetching all records
// 		var offset int
// 		if recordSize != -1 {
// 			offset = (page - 1) * recordSize
// 		}

// 		// Call the AdminService to get the Admins from the database
// 		return m.AdminService.GetAllAdmins(recordSize, offset, gender)
// 	}
// 	if flag == "false" {
// 		// If page is -1, set recordSize to -1 to fetch all records
// 		if page == -1 {
// 			recordSize = -1
// 		}

// 		// Calculate offset if not fetching all records
// 		var offset int
// 		if recordSize != -1 {
// 			offset = (page - 1) * recordSize
// 		}

// 		// Call the AdminService to get the Admins from the database
// 		return services.GetAllSUsers(recordSize, offset, gender)
// 	}
// 	return []response.Response{}, error
// }

func (m *AdminManager) GetAllAdmins(u *UserManager, page int, recordSize int, gender string, flag string) ([]response.Response, error) {
	// Normalize recordSize for fetching all records
	if page == -1 {
		recordSize = -1
	}

	// Calculate offset for pagination
	var offset int
	if recordSize != -1 {
		if page <= 0 {
			return nil, fmt.Errorf("invalid page number: %d", page)
		}
		offset = (page - 1) * recordSize
	}

	// Route to appropriate service based on flag
	switch flag {
	case "true":
		return m.AdminService.GetAllAdmins(recordSize, offset, gender)
	case "false":
		return u.UserService.GetAllSUsers(recordSize, offset, gender)
	default:
		return nil, fmt.Errorf("invalid flag value: %s", flag)
	}
}

// GetAdminByID retrieves a single Admin by ID
func (m *AdminManager) GetAdminByID(u *UserManager, id int, flag string) (*response.Response, error) {
	if flag == "true" {
		return m.AdminService.GetAdminByID(id)
	} else {
		nextID, err := services.GetNextSequence("users") // "users" collection or "admins"
		if err != nil {
			return nil, fmt.Errorf("error getting next sequence: %v", err)
		}
		idStr := strconv.Itoa(nextID)
		return u.UserService.GetUserByID(idStr)
	}
}

// UpdateAdmin modifies an existing Admin by calling the service
func (m *AdminManager) UpdateAdmin(Admin *request.Request) error {
	return m.AdminService.UpdateAdmin(Admin)
}

// DeleteAdmin removes a Admin by ID by calling the service
func (m *AdminManager) DeleteAdmin(id int) error {
	return m.AdminService.DeleteAdmin(id)
}
