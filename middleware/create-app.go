package middleware

import "github.com/gofiber/fiber/v2"

func CreateApp() *fiber.App {

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 mb limit upload files
	})

	return app
}
