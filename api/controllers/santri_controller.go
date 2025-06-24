package controllers

import (
	"siakad/api/models"
	"siakad/config"
	"siakad/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllSantri(c *fiber.Ctx) error {
	sekolahID, ok := c.Locals("sekolah_id").(uint)
	if !ok || sekolahID == 0 {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	var santri []models.Santri
	err := config.DB.
		Where("sekolah_id = ?", sekolahID).
		Find(&santri).Error
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengambil data santri")
	}

	return c.JSON(santri)
}
