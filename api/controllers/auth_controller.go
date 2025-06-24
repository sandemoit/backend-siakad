package controllers

import (
	"siakad/api/dto"
	"siakad/api/models"
	"siakad/api/service"
	"siakad/config"
	"siakad/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Permintaan tidak valid"})
	}

	user, err := service.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "Kredensial tidak valid")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal menghasilkan Token")
	}

	utils.SetCookie(c, "token", token, 24*60*60, true)

	return c.JSON(fiber.Map{
		"message": "Login Berhasil",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func Register(c *fiber.Ctx) error {
	var req dto.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal memulai transaksi")
	}

	user := models.User{
		UUID:     utils.GenerateUID(),
		Email:    req.Email,
		Password: utils.GeneratePassword(req.Password),
		Role:     req.Role,
		Name:     req.Name,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal menyimpan data")
	}

	return utils.ResponseSuccess(c, fiber.StatusCreated, "Akun Anda berhasil dibuat")
}

func Logout(c *fiber.Ctx) error {
	// Hapus token dari cookie
	utils.RevokeCookie(c, "token")

	return utils.ResponseSuccess(c, fiber.StatusOK, "Logout Berhasil")
}
