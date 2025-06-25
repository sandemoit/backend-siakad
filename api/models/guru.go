package models

import (
	"time"

	"gorm.io/gorm"
)

type Guru struct {
	ID             uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID           string         `json:"uuid" gorm:"type:char(36);uniqueIndex;not null"`
	UserID         uint           `json:"user_id" gorm:"not null;uniqueIndex"`
	Nip            string         `json:"nip" gorm:"type:varchar(20);uniqueIndex;not null"`
	SekolahID      uint           `gorm:"null;index" json:"sekolah_id"`
	NamaLengkap    string         `json:"nama_lengkap" gorm:"type:varchar(100);not null"`
	JenisKelamin   string         `json:"jenis_kelamin" gorm:"type:varchar(20);not null"`
	TanggalLahir   time.Time      `json:"tanggal_lahir" gorm:"type:date;not null"`
	TempatLahir    string         `json:"tempat_lahir" gorm:"type:varchar(50);not null"`
	Alamat         string         `json:"alamat" gorm:"type:text;not null"`
	Telepon        string         `json:"telepon" gorm:"type:varchar(15);not null"`
	Email          string         `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Jabatan        string         `json:"jabatan" gorm:"type:varchar(50);not null"`
	BidangKeahlian string         `json:"bidang_keahlian" gorm:"type:varchar(50);not null"`
	TanggalGabung  time.Time      `json:"tanggal_gabung" gorm:"type:date;not null"`
	Status         string         `json:"status" gorm:"type:varchar(10);default:'aktif';not null"`
	Gaji           float64        `json:"gaji" gorm:"type:decimal(10,2);not null"`
	Foto           *string        `json:"foto,omitempty" gorm:"type:varchar(255)"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Sekolah Sekolah `gorm:"foreignKey:SekolahID;constraint:OnDelete:SET NULL;" json:"sekolah,omitempty"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user,omitempty"`
}
