package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	UUID      string         `gorm:"uniqueIndex;size:36;not null"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Role      string         `gorm:"size:50;not null" json:"role"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
