package routes

import (
	"siakad/api/controllers"
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func SantriRoute(api fiber.Router) {
	// authentication routes
	api.Get("/santri", middleware.RoleGuard("ustadz", "admin"), controllers.GetAllSantri)
}
