package route

import (
	"bookstore/controller"

	"github.com/labstack/echo/v4"
)

// SetupRoutes defines all the routes for the movie API
func SetupRoutes(e *echo.Echo) {
	e.POST("/admin", controller.CreateAdmin)
	e.GET("/admin", controller.GetAllAdmins)
	e.GET("/admin/:id", controller.GetAdmin)
	e.PUT("/admin/:id", controller.UpdateAdmin)
	e.DELETE("/admin/:id", controller.DeleteAdmin)
}
