package repository

import (
	"siakad/api/models"

	"gorm.io/gorm"
)

type SantriRepository interface {
	FindAllBySekolahID(db *gorm.DB, sekolahID uint) ([]models.Santri, error)
	FindByIDAndSekolahID(db *gorm.DB, id, sekolahID uint) (models.Santri, error)
	FindDeletedByIDAndSekolahID(db *gorm.DB, id, sekolahID uint) (models.Santri, error)
	CheckNISExists(db *gorm.DB, nis string, sekolahID, excludeID uint) (bool, error)
	Create(db *gorm.DB, santri models.Santri) (models.Santri, error)
	Update(db *gorm.DB, santri models.Santri) (models.Santri, error)
	Delete(db *gorm.DB, id, sekolahID uint) error
	Restore(db *gorm.DB, id, sekolahID uint) error
	BulkDelete(db *gorm.DB, ids []uint, sekolahID uint) (int, error)
}

type santriRepository struct{}

func NewSantriRepository() SantriRepository {
	return &santriRepository{}
}

func (r *santriRepository) FindAllBySekolahID(db *gorm.DB, sekolahID uint) ([]models.Santri, error) {
	var santri []models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("sekolah_id = ?", sekolahID).
		Order("created_at DESC").
		Find(&santri).Error
	return santri, err
}

func (r *santriRepository) FindByIDAndSekolahID(db *gorm.DB, id, sekolahID uint) (models.Santri, error) {
	var santri models.Santri
	err := db.Where("id = ? AND sekolah_id = ?", id, sekolahID).
		First(&santri).Error
	return santri, err
}

func (r *santriRepository) FindDeletedByIDAndSekolahID(db *gorm.DB, id, sekolahID uint) (models.Santri, error) {
	var santri models.Santri
	err := db.Unscoped().
		Preload("Sekolah").
		Preload("User").
		Where("id = ? AND sekolah_id = ? AND deleted_at IS NOT NULL", id, sekolahID).
		First(&santri).Error
	return santri, err
}

func (r *santriRepository) CheckNISExists(db *gorm.DB, nis string, sekolahID, excludeID uint) (bool, error) {
	var count int64
	query := db.Model(&models.Santri{}).
		Where("nis = ? AND sekolah_id = ?", nis, sekolahID)

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	err := query.Count(&count).Error
	return count > 0, err
}

func (r *santriRepository) Create(db *gorm.DB, santri models.Santri) (models.Santri, error) {
	err := db.Create(&santri).Error
	if err != nil {
		return santri, err
	}
	return santri, nil
}

func (r *santriRepository) Update(db *gorm.DB, santri models.Santri) (models.Santri, error) {
	err := db.Save(&santri).Error
	if err != nil {
		return santri, err
	}
	return santri, nil
}

func (r *santriRepository) Delete(db *gorm.DB, id, sekolahID uint) error {
	return db.Where("id = ? AND sekolah_id = ?", id, sekolahID).
		Delete(&models.Santri{}).Error
}

func (r *santriRepository) Restore(db *gorm.DB, id, sekolahID uint) error {
	return db.Unscoped().
		Model(&models.Santri{}).
		Where("id = ? AND sekolah_id = ?", id, sekolahID).
		Update("deleted_at", nil).Error
}

func (r *santriRepository) BulkDelete(db *gorm.DB, ids []uint, sekolahID uint) (int, error) {
	result := db.Where("id IN ? AND sekolah_id = ?", ids, sekolahID).
		Delete(&models.Santri{})

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

// Additional helper methods for more complex queries

func (r *santriRepository) FindByNIS(db *gorm.DB, nis string, sekolahID uint) (models.Santri, error) {
	var santri models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("nis = ? AND sekolah_id = ?", nis, sekolahID).
		First(&santri).Error
	return santri, err
}

func (r *santriRepository) FindByKelasID(db *gorm.DB, kelasID, sekolahID uint) ([]models.Santri, error) {
	var santri []models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("kelas_id = ? AND sekolah_id = ?", kelasID, sekolahID).
		Order("nama_lengkap ASC").
		Find(&santri).Error
	return santri, err
}

func (r *santriRepository) FindByAsramaID(db *gorm.DB, asramaID, sekolahID uint) ([]models.Santri, error) {
	var santri []models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("asrama_id = ? AND sekolah_id = ?", asramaID, sekolahID).
		Order("nama_lengkap ASC").
		Find(&santri).Error
	return santri, err
}

func (r *santriRepository) FindByStatus(db *gorm.DB, status string, sekolahID uint) ([]models.Santri, error) {
	var santri []models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("status = ? AND sekolah_id = ?", status, sekolahID).
		Order("nama_lengkap ASC").
		Find(&santri).Error
	return santri, err
}

func (r *santriRepository) FindActiveBySekolahID(db *gorm.DB, sekolahID uint) ([]models.Santri, error) {
	var santri []models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("sekolah_id = ? AND is_active = ?", sekolahID, true).
		Order("nama_lengkap ASC").
		Find(&santri).Error
	return santri, err
}

func (r *santriRepository) CountBySekolahID(db *gorm.DB, sekolahID uint) (int64, error) {
	var count int64
	err := db.Model(&models.Santri{}).
		Where("sekolah_id = ?", sekolahID).
		Count(&count).Error
	return count, err
}

func (r *santriRepository) CountActiveBySekolahID(db *gorm.DB, sekolahID uint) (int64, error) {
	var count int64
	err := db.Model(&models.Santri{}).
		Where("sekolah_id = ? AND is_active = ?", sekolahID, true).
		Count(&count).Error
	return count, err
}

func (r *santriRepository) SearchByName(db *gorm.DB, name string, sekolahID uint) ([]models.Santri, error) {
	var santri []models.Santri
	err := db.Preload("Sekolah").
		Preload("User").
		Where("nama_lengkap ILIKE ? AND sekolah_id = ?", "%"+name+"%", sekolahID).
		Order("nama_lengkap ASC").
		Find(&santri).Error
	return santri, err
}

func (r *santriRepository) FindPaginated(db *gorm.DB, sekolahID uint, offset, limit int) ([]models.Santri, int64, error) {
	var santri []models.Santri
	var total int64

	// Count total records
	err := db.Model(&models.Santri{}).
		Where("sekolah_id = ?", sekolahID).
		Count(&total).Error
	if err != nil {
		return santri, 0, err
	}

	// Get paginated records
	err = db.Preload("Sekolah").
		Preload("User").
		Where("sekolah_id = ?", sekolahID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&santri).Error

	return santri, total, err
}
