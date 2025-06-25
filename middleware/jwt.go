package middleware

import (
	"os"
	"siakad/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("token")
		if tokenStr == "" {
			authHeader := c.Get("Authorization")
			if authHeader != "" {
				tokenStr = authHeader
			}
		}

		if tokenStr == "" {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Unauthorized: Token tidak ditemukan")
		}

		token, err := utils.VerifyToken(tokenStr)
		if err != nil || token == nil || !token.Valid {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid: "+err.Error())
		}

		// Optional: Check expired (karena Parse sudah otomatis cek exp juga)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if exp, ok := claims["exp"].(float64); ok {
				if int64(exp) < time.Now().Unix() {
					return utils.ResponseError(c, fiber.StatusUnauthorized, "Token telah kadaluarsa")
				}
			}
		}

		c.Locals("user", token)
		return c.Next()
	}
}

func GuestOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("token")
		if tokenStr == "" {
			return c.Next()
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Next()
		}

		return utils.ResponseError(c, fiber.StatusForbidden, "Anda sudah login")
	}
}

func RoleGuard(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken, ok := c.Locals("user").(*jwt.Token)
		if !ok || userToken == nil {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid")
		}

		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Klaim token tidak valid")
		}

		roleVal, ok := claims["role"]
		if !ok {
			return utils.ResponseError(c, fiber.StatusForbidden, "Peran tidak ditemukan")
		}

		role, ok := roleVal.(string)
		if !ok {
			return utils.ResponseError(c, fiber.StatusForbidden, "Tipe peran tidak valid")
		}

		for _, r := range allowedRoles {
			if role == r {
				return c.Next()
			}
		}

		return utils.ResponseError(c, fiber.StatusForbidden, "Akses ditolak untuk peran "+role)
	}
}
