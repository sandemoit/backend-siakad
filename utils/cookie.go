package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, name, value string, maxAge int) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
	})
}

func RevokeCookie(c *fiber.Ctx, name string) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HTTPOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
	})
}
