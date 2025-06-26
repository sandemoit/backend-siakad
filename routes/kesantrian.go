package routes

import (
	"siakad/api/controllers"
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func KesantrianRoute(app fiber.Router) {
	santri := app.Group("/santri")

	// Middleware authentication untuk semua route santri
	santri.Use(middleware.RoleGuard("admin", "wali_kelas"))

	// CRUD Routes
	santri.Get("/", controllers.GetAllSantri)
	santri.Get("/:id", controllers.GetSantriByID)
	santri.Post("/", controllers.CreateSantri)
	santri.Put("/:id", controllers.UpdateSantri)
	santri.Delete("/:id", controllers.DeleteSantri)

	// Additional Routes
	santri.Post("/:id/restore", controllers.RestoreSantri)
	santri.Post("/bulk-delete", controllers.BulkDeleteSantri)
}
