package middleware

import (
	"siakad/api/models"
	"siakad/config"
	"siakad/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.Locals("user").(*jwt.Token)

		if !ok || token == nil {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Klaim token tidak valid")
		}

		sekolahID, ok := claims["sekolah_id"]
		if !ok {
			return utils.ResponseError(c, fiber.StatusForbidden, "Akses ditolak: sekolah_id tidak ditemukan di token")
		}

		// Cek apakah tenant aktif
		var tenant models.Tenant
		if err := config.DB.Where("sekolah_id = ? AND is_active = ?", sekolahID, true).First(&tenant).Error; err != nil {
			return utils.ResponseError(c, fiber.StatusForbidden, "Akses ditolak: Tenant tidak aktif atau tidak ditemukan")
		}

		// Simpan ke context
		c.Locals("sekolah_id", sekolahID)

		return c.Next()
	}
}
