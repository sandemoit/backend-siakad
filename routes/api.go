package routes

import (
	"siakad/middleware"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(api fiber.Router) {
	AuthRoute(api)

	protectedSantri := api.Group("", middleware.JWTProtected(), middleware.TenantMiddleware())
	KesantrianRoute(protectedSantri)
}
