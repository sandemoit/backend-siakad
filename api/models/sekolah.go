package models

import (
	"time"

	"gorm.io/gorm"
)

type Sekolah struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UUID      string         `gorm:"uniqueIndex;size:36;not null" json:"uuid"`
	Nama      string         `gorm:"size:100;not null" json:"nama"`
	Provinsi  string         `gorm:"size:50;not null" json:"provinsi"`
	Kabupaten string         `gorm:"size:50;not null" json:"kabupaten"`
	Kecamatan string         `gorm:"size:50;not null" json:"kecamatan"`
	Alamat    string         `gorm:"size:255;not null" json:"alamat"`
	NPSN      string         `gorm:"size:20;not null" json:"npsn"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Telepon   string         `gorm:"size:20;not null" json:"telepon"`
	Logo      string         `gorm:"size:255;" json:"logo"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Santris []Santri `gorm:"foreignKey:SekolahID" json:"santris,omitempty"`
	Users   []User   `gorm:"foreignKey:SekolahID" json:"users,omitempty"`
}
