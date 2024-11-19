// manager/admin_manager.go

package manager

import (
	"bookstore/request"
	"bookstore/response"
	"bookstore/services"
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

// CreateAdmin adds a new Admin by interacting with the service
func (m *AdminManager) CreateAdmin(Admin *request.Request) (*response.Response, error) {
	log.Println("Mgr begin")

	// Call the service to create the Admin
	res, err := m.AdminService.CreateAdmin(Admin)
	if err != nil {
		return nil, err
	}

	log.Println("Mgr close")
	return res, nil
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

func (m *AdminManager) GetAllAdmins(page, recordSize int, genre string) ([]response.Response, error) {
	// If page is -1, set recordSize to -1 to fetch all records
	if page == -1 {
		recordSize = -1
	}

	// Calculate offset if not fetching all records
	var offset int
	if recordSize != -1 {
		offset = (page - 1) * recordSize
	}

	// Call the AdminService to get the Admins from the database
	return m.AdminService.GetAllAdmins(recordSize, offset, genre)
}

// GetAdminByID retrieves a single Admin by ID
func (m *AdminManager) GetAdminByID(id int) (*response.Response, error) {
	return m.AdminService.GetAdminByID(id)
}

// UpdateAdmin modifies an existing Admin by calling the service
func (m *AdminManager) UpdateAdmin(Admin *request.Request) error {
	return m.AdminService.UpdateAdmin(Admin)
}

// DeleteAdmin removes a Admin by ID by calling the service
func (m *AdminManager) DeleteAdmin(id int) error {
	return m.AdminService.DeleteAdmin(id)
}
