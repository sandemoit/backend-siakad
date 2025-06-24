package routes

import (
	"siakad/api/controllers"
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	api.Post("/auth/login", controllers.Login)
	api.Post("/auth/register", controllers.Register)

	api.Post("/auth/logout", middleware.JWTProtected(), controllers.Logout)

	// penggunaan role
	// api.Post("/auth/login", middleware.JWTProtected(), middleware.RoleGuard("admin"), controllers.Login)

	// api.Get("/me", middleware.JWTProtected(), func(c *fiber.Ctx) error {
	// 	user := c.Locals("user")
	// 	return c.JSON(user)
	// })
}
