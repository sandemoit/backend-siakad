package controllers

import (
	"fmt"
	"siakad/api/dto"
	"siakad/api/models"
	"siakad/api/service"
	"siakad/config"
	"siakad/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
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

	// return c.JSON(fiber.Map{
	// 	"message": "Login Berhasil",
	// 	"user": fiber.Map{
	// 		"id":    user.ID,
	// 		"name":  user.Name,
	// 		"email": user.Email,
	// 		"role":  user.Role,
	// 	},
	// })

	return utils.ResponseSuccess(c, fiber.StatusOK, "Login Berhasil")
}

func Register(c *fiber.Ctx) error {
	var req dto.RegisterSekolahRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal memulai transaksi")
	}

	// Simpan sekolah dulu
	sekolah := models.Sekolah{
		UUID:      utils.GenerateUID(),
		Nama:      req.NamaSekolah,
		Provinsi:  req.Provinsi,
		Kabupaten: req.Kabupaten,
		Kecamatan: req.Kecamatan,
		Alamat:    req.Alamat,
		NPSN:      req.NPSN,
		Email:     req.EmailSekolah,
		Telepon:   req.Telepon,
		Logo:      req.Logo,
	}
	if err := tx.Create(&sekolah).Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	// Autogenerate password 8 karakter
	genPassword := utils.GenerateRandomPassword(8)

	// Simpan user admin sekolah
	user := models.User{
		UUID:      utils.GenerateUID(),
		Name:      req.NamaAdmin,
		Email:     req.EmailAdmin,
		Password:  utils.GeneratePassword(genPassword), // hash password
		Role:      "admin",
		SekolahID: sekolah.ID,
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	guru := models.Guru{
		UUID:           utils.GenerateUID(),
		UserID:         user.ID,
		Nip:            "000000000", // bisa kamu buat random generator NIP atau kosongin dulu
		SekolahID:      sekolah.ID,
		NamaLengkap:    req.NamaAdmin,
		JenisKelamin:   "-",        // pastikan dari req juga
		TanggalLahir:   time.Now(), // atau dari req
		TempatLahir:    "-",
		Alamat:         req.Alamat,
		Telepon:        req.Telepon,
		Email:          req.EmailAdmin,
		Jabatan:        "Admin Pondok",
		BidangKeahlian: "-",
		TanggalGabung:  time.Now(),
		Status:         "aktif",
		Gaji:           0,
	}
	if err := tx.Create(&guru).Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	invoice := models.Invoice{
		NoInvoice:     utils.GenerateInvoiceNumber(), // contoh: INV-202406230001
		Tanggal:       time.Now(),
		StatusPayment: "unpaid",
	}
	if err := tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	tenant := models.Tenant{
		SekolahID: sekolah.ID,
		InvoiceID: invoice.ID,
		Domain:    fmt.Sprintf("%s.siakadpesantren.id", utils.Slugify(sekolah.Nama)),
		IsActive:  true,
	}
	if err := tx.Create(&tenant).Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal menyimpan data")
	}

	data := fiber.Map{
		"email":    user.Email,
		"password": genPassword, // optional ditampilkan
	}

	// Balikin password-nya ke user kalau perlu
	return utils.ResponseSuccess(c, fiber.StatusCreated, "Berhasil membuat akun sekolah", data)
}

func Logout(c *fiber.Ctx) error {
	// Hapus token dari cookie
	utils.RevokeCookie(c, "token")

	return utils.ResponseSuccess(c, fiber.StatusOK, "Logout Berhasil")
}
