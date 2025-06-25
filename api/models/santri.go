package models

import (
	"time"

	"gorm.io/gorm"
)

type Santri struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `gorm:"index" json:"user_id"`
	NIS           string         `gorm:"unique" json:"nis"`
	NamaLengkap   string         `gorm:"size:100;not null" json:"nama_lengkap"`
	JenisKelamin  string         `gorm:"size:1" json:"jenis_kelamin"`
	TanggalLahir  *time.Time     `json:"tanggal_lahir"`
	TempatLahir   string         `json:"tempat_lahir"`
	Alamat        string         `gorm:"type:text" json:"alamat"`
	Telepon       string         `json:"telepon"`
	NamaAyah      string         `json:"nama_ayah"`
	NamaIbu       string         `json:"nama_ibu"`
	PekerjaanAyah string         `json:"pekerjaan_ayah"`
	PekerjaanIbu  string         `json:"pekerjaan_ibu"`
	TeleponOrtu   string         `json:"telepon_ortu"`
	TanggalMasuk  *time.Time     `json:"tanggal_masuk"`
	Status        string         `gorm:"size:20" json:"status"`
	KelasID       uint           `json:"kelas_id"`
	AsramaID      uint           `json:"asrama_id"`
	Foto          string         `json:"foto"`
	SekolahID     uint           `gorm:"index" json:"sekolah_id"`
	IsActive      bool           `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Sekolah Sekolah `gorm:"foreignKey:SekolahID" json:"sekolah,omitempty"`
	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
