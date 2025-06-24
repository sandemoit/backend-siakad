package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UUID      string         `gorm:"uniqueIndex;size:36;not null" json:"uuid"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Role      string         `gorm:"type:varchar(32);not null" json:"role"`
	SekolahID uint           `gorm:"null;index" json:"sekolah_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Sekolah Sekolah `gorm:"foreignKey:SekolahID;constraint:OnDelete:CASCADE;" json:"sekolah"`
}
