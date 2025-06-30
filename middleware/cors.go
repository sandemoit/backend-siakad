package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(app *fiber.App) error {
	var allowedOrigins []string

	if os.Getenv("MODE") == "dev" {
		origins := strings.Split(os.Getenv("ALLOWED_ORIGINS_DEV"), ",")
		for _, origin := range origins {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		}
	} else {
		origins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
		for _, origin := range origins {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		}
	}

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {

			for _, o := range allowedOrigins {
				if o == origin {
					return true
				}
			}
			return false
		},
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Set-Cookie",
		AllowCredentials: true,
	}))

	return nil
}
