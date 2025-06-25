package routes

import (
	"siakad/api/controllers"
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	// authentication routes
	api.Post("/auth/login", middleware.GuestOnly(), controllers.Login)
	api.Post("/auth/register", middleware.GuestOnly(), controllers.Register)

	api.Post("/auth/logout", controllers.Logout)
}
