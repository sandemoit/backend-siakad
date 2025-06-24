package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, name, value string, maxAge int, httpOnly bool) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		HTTPOnly: httpOnly, // Prevents JavaScript access
		Secure:   false,    // Set to true if using HTTPS
		SameSite: "Lax",    // Adjust based on your needs
		Path:     "/",      // Cookie path
	})
}

func RevokeCookie(c *fiber.Ctx, name string) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Path:     "/",
	})
}
