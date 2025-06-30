package routes

import (
	"siakad/api/controllers"
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	// authentication routes
	auth := api.Group("/auth")

	auth.Post("/login", middleware.GuestOnly(), controllers.Login)
	auth.Post("/register", middleware.GuestOnly(), controllers.Register)

	auth.Get("/verify", middleware.JWTProtected(), middleware.TenantMiddleware(), controllers.VerifyToken)
	auth.Post("/refresh", controllers.RefreshToken)

	auth.Use(middleware.JWTProtected(), middleware.TenantMiddleware(), middleware.RoleGuard("admin", "wali_kelas"))
	auth.Post("/logout", controllers.Logout)
}
