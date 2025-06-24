package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PublicRoute(api fiber.Router) {
	AuthRoute(api)
	SantriRoute(api)
}
