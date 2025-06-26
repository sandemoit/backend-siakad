package service

import (
	"errors"
	"siakad/api/dto"
	"siakad/api/models"
	"siakad/api/repository"
	"siakad/config"

	"gorm.io/gorm"
)

type SantriService interface {
	GetAllSantri(sekolahID uint) ([]models.Santri, error)
	GetSantriByID(id, sekolahID uint) (models.Santri, error)
	CreateSantri(req dto.SantriRequest) (models.Santri, error)
	UpdateSantri(id uint, req dto.SantriRequest, sekolahID uint) (models.Santri, error)
	DeleteSantri(id, sekolahID uint) error
	RestoreSantri(id, sekolahID uint) (models.Santri, error)
	BulkDeleteSantri(ids []uint, sekolahID uint) (int, error)
}

type santriService struct {
	repo repository.SantriRepository
}

func NewSantriService(repo repository.SantriRepository) SantriService {
	return &santriService{repo}
}

func (s *santriService) GetAllSantri(sekolahID uint) ([]models.Santri, error) {
	return s.repo.FindAllBySekolahID(config.DB, sekolahID)
}

func (s *santriService) GetSantriByID(id, sekolahID uint) (models.Santri, error) {
	santri, err := s.repo.FindByIDAndSekolahID(config.DB, id, sekolahID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return santri, errors.New("santri not found")
		}
		return santri, err
	}
	return santri, nil
}

func (s *santriService) CreateSantri(req dto.SantriRequest) (models.Santri, error) {
	// Begin transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		return models.Santri{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if NIS already exists
	exists, err := s.repo.CheckNISExists(tx, req.NIS, req.SekolahID, 0)
	if err != nil {
		tx.Rollback()
		return models.Santri{}, err
	}
	if exists {
		tx.Rollback()
		return models.Santri{}, errors.New("NIS sudah digunakan")
	}

	// Convert DTO to model
	santri := models.Santri{
		UserID:        req.UserID,
		NIS:           req.NIS,
		NamaLengkap:   req.NamaLengkap,
		JenisKelamin:  req.JenisKelamin,
		TanggalLahir:  req.TanggalLahir,
		TempatLahir:   req.TempatLahir,
		Alamat:        req.Alamat,
		Telepon:       req.Telepon,
		NamaAyah:      req.NamaAyah,
		NamaIbu:       req.NamaIbu,
		PekerjaanAyah: req.PekerjaanAyah,
		PekerjaanIbu:  req.PekerjaanIbu,
		TeleponOrtu:   req.TeleponOrtu,
		TanggalMasuk:  req.TanggalMasuk,
		Status:        req.Status,
		KelasID:       req.KelasID,
		AsramaID:      req.AsramaID,
		Foto:          req.Foto,
		SekolahID:     req.SekolahID,
		IsActive:      req.IsActive,
	}

	// Create santri
	createdSantri, err := s.repo.Create(tx, santri)
	if err != nil {
		tx.Rollback()
		return models.Santri{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.Santri{}, err
	}

	// Get complete data with relations
	return s.repo.FindByIDAndSekolahID(config.DB, createdSantri.ID, req.SekolahID)
}

func (s *santriService) UpdateSantri(id uint, req dto.SantriRequest, sekolahID uint) (models.Santri, error) {
	// Begin transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		return models.Santri{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if santri exists
	existingSantri, err := s.repo.FindByIDAndSekolahID(tx, id, sekolahID)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Santri{}, errors.New("santri not found")
		}
		return models.Santri{}, err
	}

	// Check if NIS already exists (exclude current santri)
	exists, err := s.repo.CheckNISExists(tx, req.NIS, req.SekolahID, id)
	if err != nil {
		tx.Rollback()
		return models.Santri{}, err
	}
	if exists {
		tx.Rollback()
		return models.Santri{}, errors.New("NIS sudah digunakan")
	}

	// Update fields
	existingSantri.UserID = req.UserID
	existingSantri.NIS = req.NIS
	existingSantri.NamaLengkap = req.NamaLengkap
	existingSantri.JenisKelamin = req.JenisKelamin
	existingSantri.TanggalLahir = req.TanggalLahir
	existingSantri.TempatLahir = req.TempatLahir
	existingSantri.Alamat = req.Alamat
	existingSantri.Telepon = req.Telepon
	existingSantri.NamaAyah = req.NamaAyah
	existingSantri.NamaIbu = req.NamaIbu
	existingSantri.PekerjaanAyah = req.PekerjaanAyah
	existingSantri.PekerjaanIbu = req.PekerjaanIbu
	existingSantri.TeleponOrtu = req.TeleponOrtu
	existingSantri.TanggalMasuk = req.TanggalMasuk
	existingSantri.Status = req.Status
	existingSantri.KelasID = req.KelasID
	existingSantri.AsramaID = req.AsramaID
	existingSantri.Foto = req.Foto
	existingSantri.IsActive = req.IsActive

	// Update santri
	updatedSantri, err := s.repo.Update(tx, existingSantri)
	if err != nil {
		tx.Rollback()
		return models.Santri{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.Santri{}, err
	}

	// Get complete data with relations
	return s.repo.FindByIDAndSekolahID(config.DB, updatedSantri.ID, sekolahID)
}

func (s *santriService) DeleteSantri(id, sekolahID uint) error {
	// Begin transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if santri exists
	_, err := s.repo.FindByIDAndSekolahID(tx, id, sekolahID)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("santri not found")
		}
		return err
	}

	// Soft delete santri
	err = s.repo.Delete(tx, id, sekolahID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

func (s *santriService) RestoreSantri(id, sekolahID uint) (models.Santri, error) {
	// Begin transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		return models.Santri{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if deleted santri exists
	_, err := s.repo.FindDeletedByIDAndSekolahID(tx, id, sekolahID)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Santri{}, errors.New("santri not found")
		}
		return models.Santri{}, err
	}

	// Restore santri
	err = s.repo.Restore(tx, id, sekolahID)
	if err != nil {
		tx.Rollback()
		return models.Santri{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.Santri{}, err
	}

	// Get restored data
	return s.repo.FindByIDAndSekolahID(config.DB, id, sekolahID)
}

func (s *santriService) BulkDeleteSantri(ids []uint, sekolahID uint) (int, error) {
	// Begin transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Bulk delete santri
	deletedCount, err := s.repo.BulkDelete(tx, ids, sekolahID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return deletedCount, nil
}
