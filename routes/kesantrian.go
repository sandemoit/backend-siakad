package routes

import (
	"siakad/api/controllers"
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func KesantrianRoute(app fiber.Router) {
	// Group routes dengan prefix /api/santri
	api := app.Group("/santri")

	// Middleware authentication untuk semua route santri
	api.Use(middleware.RoleGuard("admin", "wali_kelas"))

	// CRUD Routes
	api.Get("/", controllers.GetAllSantri)       // GET /api/santri
	api.Get("/:id", controllers.GetSantriByID)   // GET /api/santri/:id
	api.Post("/", controllers.CreateSantri)      // POST /api/santri
	api.Put("/:id", controllers.UpdateSantri)    // PUT /api/santri/:id
	api.Delete("/:id", controllers.DeleteSantri) // DELETE /api/santri/:id

	// Additional Routes
	api.Post("/:id/restore", controllers.RestoreSantri)    // POST /api/santri/:id/restore
	api.Post("/bulk-delete", controllers.BulkDeleteSantri) // POST /api/santri/bulk-delete
}
