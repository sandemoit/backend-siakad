package routes

import (
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// CORS
	middleware.Cors(app)

	// Logger API
	middleware.LogAPI(app)

	// Definisi publicRoutes yang tidak perlu autentikasi
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusForbidden).SendString("Access Denied")
	})

	api := app.Group("/api/v1")
	// api.Use(middleware.CheckSession)

	// Public routes
	PublicRoute(api)
}
