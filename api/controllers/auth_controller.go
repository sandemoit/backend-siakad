package controllers

import (
	"siakad/api/models"
	"siakad/api/service"
	"siakad/config"
	"siakad/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Name     string `json:"name,omitempty"` // Optional, can be used for registration
	Role     string `json:"role,omitempty"` // Optional, can be used for registration
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Permintaan tidak valid"})
	}

	user, err := service.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Kredensial tidak valid"})
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghasilkan Token"})
	}

	// 🔐 Set token ke dalam Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,  // ❗ Tidak bisa diakses dari JS
		Secure:   false, // Aktifkan kalau pakai HTTPS
		SameSite: "Lax", // Lax/SameSiteNone/Strict sesuai kebutuhan
		Path:     "/",
	})

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
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal memulai transaksi")
	}

	user := models.User{
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
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "Logout Berhasil",
	})
}
