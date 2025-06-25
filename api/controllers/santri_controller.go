package controllers

import (
	"siakad/api/models"
	"siakad/config"
	"siakad/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahFloat, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}
	sekolahID := uint(sekolahFloat)

	var santri []models.Santri
	err := config.DB.
		Preload("Sekolah").
		Where("sekolah_id = ?", sekolahID).
		Find(&santri).Error
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengambil data santri")
	}

	return c.JSON(santri)
}
