package controllers

import (
	"fmt"
	"siakad/api/dto"
	"siakad/api/repository"
	"siakad/api/service"
	"siakad/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var santriService = service.NewSantriService(
	repository.NewSantriRepository(),
)

// GetAllSantri - Mengambil semua data santri berdasarkan sekolah
func GetAllSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	data, err := santriService.GetAllSantri(uint(sekolahID))
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengambil data santri")
	}

	return utils.ResponGetData(c, fiber.StatusOK, data)
}

// GetSantriByID - Mengambil data santri berdasarkan ID
func GetSantriByID(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID santri tidak valid")
	}

	data, err := santriService.GetSantriByID(uint(id), uint(sekolahID))
	if err != nil {
		if err.Error() == "santri not found" {
			return utils.ResponseError(c, fiber.StatusNotFound, "Data santri tidak ditemukan")
		}
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengambil data santri")
	}

	return utils.ResponGetData(c, fiber.StatusOK, data)
}

// CreateSantri - Membuat data santri baru
func CreateSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	var req dto.SantriRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Format data tidak valid")
	}

	// Set sekolah ID dari context
	req.SekolahID = uint(sekolahID)

	// Validasi data
	if err := utils.ValidateStruct(req); err != nil {
		message := fmt.Sprintf("%v", err)
		return utils.ResponseError(c, fiber.StatusBadRequest, message)
	}

	data, err := santriService.CreateSantri(req)
	if err != nil {
		if err.Error() == "NIS sudah digunakan" {
			return utils.ResponseError(c, fiber.StatusConflict, "NIS sudah digunakan")
		}
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal membuat data santri")
	}

	return utils.ResponseSuccess(c, fiber.StatusCreated, "Data santri berhasil dibuat", data)
}

// UpdateSantri - Mengupdate data santri
func UpdateSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID santri tidak valid")
	}

	var req dto.SantriRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Format data tidak valid")
	}

	// Set sekolah ID dari context
	req.SekolahID = uint(sekolahID)

	// Validasi data
	if err := utils.ValidateStruct(req); err != nil {
		message := fmt.Sprintf("%v", err)
		return utils.ResponseError(c, fiber.StatusBadRequest, message)
	}

	data, err := santriService.UpdateSantri(uint(id), req, uint(sekolahID))
	if err != nil {
		if err.Error() == "santri not found" {
			return utils.ResponseError(c, fiber.StatusNotFound, "Data santri tidak ditemukan")
		}
		if err.Error() == "NIS sudah digunakan" {
			return utils.ResponseError(c, fiber.StatusConflict, "NIS sudah digunakan")
		}
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengupdate data santri")
	}

	return utils.ResponseSuccess(c, fiber.StatusOK, "Data santri berhasil diupdate", data)
}

// DeleteSantri - Menghapus data santri (soft delete)
func DeleteSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID santri tidak valid")
	}

	err = santriService.DeleteSantri(uint(id), uint(sekolahID))
	if err != nil {
		if err.Error() == "santri not found" {
			return utils.ResponseError(c, fiber.StatusNotFound, "Data santri tidak ditemukan")
		}
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal menghapus data santri")
	}

	return utils.ResponseSuccess(c, fiber.StatusOK, "Data santri berhasil dihapus", nil)
}

// RestoreSantri - Mengembalikan data santri yang telah dihapus
func RestoreSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID santri tidak valid")
	}

	data, err := santriService.RestoreSantri(uint(id), uint(sekolahID))
	if err != nil {
		if err.Error() == "santri not found" {
			return utils.ResponseError(c, fiber.StatusNotFound, "Data santri tidak ditemukan")
		}
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengembalikan data santri")
	}

	return utils.ResponseSuccess(c, fiber.StatusOK, "Data santri berhasil dikembalikan", data)
}

// BulkDeleteSantri - Menghapus multiple santri sekaligus
func BulkDeleteSantri(c *fiber.Ctx) error {
	sekolahIDRaw := c.Locals("sekolah_id")
	sekolahID, ok := sekolahIDRaw.(float64)
	if !ok {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID sekolah tidak valid")
	}

	var req struct {
		IDs []uint `json:"ids" validate:"required,min=1"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Format data tidak valid")
	}

	if err := utils.ValidateStruct(req); err != nil {
		message := fmt.Sprintf("%v", err)
		return utils.ResponseError(c, fiber.StatusBadRequest, message)
	}

	deletedCount, err := santriService.BulkDeleteSantri(req.IDs, uint(sekolahID))
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal menghapus data santri")
	}

	return utils.ResponseSuccess(c, fiber.StatusOK,
		"Berhasil menghapus "+strconv.Itoa(deletedCount)+" data santri",
		map[string]int{"deleted_count": deletedCount})
}
