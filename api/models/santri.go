package models

import (
	"time"

	"gorm.io/gorm"
)

type Santri struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `gorm:"size:100;not null" json:"nama"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	SekolahID uint           `gorm:"index" json:"sekolah_id"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Sekolah Sekolah `gorm:"foreignKey:SekolahID" json:"sekolah"`
}
