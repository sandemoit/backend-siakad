package controllers

import (
	"fmt"
	"os"
	"siakad/api/dto"
	"siakad/api/models"
	"siakad/api/service"
	"siakad/config"
	"siakad/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body dto.LoginRequest true "Login Request"
// @Success 200 {object} utils.BaseResponse
// @Failure 400 {object} utils.PlatResponse
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Permintaan tidak valid"})
	}

	user, err := service.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "Kredensial tidak valid")
	}

	accessToken, err := utils.GenerateToken(user, 1*24*time.Hour)
	if err != nil {
		utils.LogError(err)
	}

	refreshToken, err := utils.GenerateToken(user, 7*24*time.Hour)
	if err != nil {
		utils.LogError(err)
	}

	utils.SetCookie(c, "access_token", accessToken, 24*60*60)
	utils.SetCookie(c, "refresh_token", refreshToken, 24*60*60)

	// return utils.ResponseSuccess(c, fiber.StatusOK, "Login Berhasil")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Berhasil",
		"user":    user,
		"token":   accessToken,
	})
}

// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body dto.RegisterSekolahRequest true "Register Request"
// @Success 200 {object} utils.BaseResponse
// @Failure 400 {object} utils.PlatResponse
// @Router /auth/register [post]
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

// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} utils.BaseResponse
// @Failure 400 {object} utils.PlatResponse
// @Router /auth/logout [post]
func Logout(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			utils.ResponseError(c, fiber.StatusInternalServerError, "Terjadi kesalahan saat logout")
		}
	}()

	// Hapus token dari cookie
	utils.ClearCookie(c, "access_token")
	utils.ClearCookie(c, "refresh_token")

	return utils.ResponseSuccess(c, fiber.StatusOK, "Logout Berhasil")
}

func VerifyToken(c *fiber.Ctx) error {
	// Ambil token dari cookie atau header Authorization
	var tokenString string

	// Cek dari cookie terlebih dahulu
	tokenString = c.Cookies("access_token")
	if tokenString == "" {
		// Jika tidak ada di cookie, cek di header Authorization
		authHeader := c.Get("Authorization")
		if authHeader != "" {
			tokenString = authHeader
		}
	}

	// Jika token masih kosong
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: Token tidak ditemukan",
		})
	}

	// Parse dan verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Algoritma signing tidak valid")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: Token tidak valid",
			"error":   err.Error(),
		})
	}

	// Cek apakah token valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: Token tidak valid",
		})
	}

	// Jika semua validasi berhasil
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Token valid",
		"user":    token.Claims.(jwt.MapClaims)["sub"], // Ambil data user dari claims
	})
}

func RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Refresh token tidak ditemukan",
		})
	}

	token, err := utils.VerifyToken(refreshToken)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Refresh token tidak valid",
		})
	}

	userID, role, err := utils.GetUserFromToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Gagal membaca token",
		})
	}

	// Bisa validasi user di DB juga jika perlu
	accessToken, _ := utils.GenerateToken(&models.User{
		ID:   userID,
		Role: role,
	}, 15*time.Minute)

	utils.SetCookie(c, "access_token", accessToken, 15*60)

	return c.JSON(fiber.Map{
		"message": "Token baru berhasil dibuat",
	})
}
