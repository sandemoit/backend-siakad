package middleware

import (
	"fmt"
	"siakad/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RoleGuard(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken, ok := c.Locals("user").(*jwt.Token)
		if !ok || userToken == nil {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Tidak sah - token tidak valid atau hilang")
		}

		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Tidak sah - klaim tidak valid")
		}

		roleVal, ok := claims["role"]
		if !ok {
			return utils.ResponseError(c, fiber.StatusForbidden, "Tidak sah - peran tidak ditemukan di token")
		}

		role, ok := roleVal.(string)
		if !ok {
			if f, ok := roleVal.(float64); ok {
				role = fmt.Sprintf("%.0f", f)
			} else {
				return utils.ResponseError(c, fiber.StatusForbidden, "Tidak sah - tipe peran tidak valid")
			}
		}

		for _, r := range allowedRoles {
			if r == role {
				return c.Next()
			}
		}

		return utils.ResponseError(c, fiber.StatusForbidden, "Tidak sah - peran tidak valid")
	}
}

func GuestOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("token")
		if tokenStr == "" {
			return c.Next() // ✅ Tidak ada token → boleh lanjut (guest)
		}

		token, err := utils.VerifyToken(tokenStr)
		if err != nil || !token.Valid {
			return c.Next() // ✅ Token rusak / invalid → tetap dianggap guest
		}

		// ❌ Token valid → berarti user sudah login
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are already authenticated",
		})
	}
}

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("token")

		// If not found in cookie, try Authorization header
		if tokenString == "" {
			authHeader := c.Get("Authorization")
			if authHeader != "" {
				tokenString = authHeader
			}
		}

		if tokenString == "" {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak ditemukan")
		}

		// Verifikasi token
		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid: "+err.Error())
		}

		// Cek tipe klaim dan pastikan token belum kadaluarsa
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid")
		}
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < utils.NowUnix() {
				return utils.ResponseError(c, fiber.StatusUnauthorized, "Token telah kadaluarsa")
			}
		}

		// Simpan token ke context
		c.Locals("user", token)
		return c.Next()
	}
}
