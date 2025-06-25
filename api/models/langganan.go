package models

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	InvoiceID uint           `json:"invoice_id" gorm:"not null;uniqueIndex"`
	SekolahID uint           `json:"sekolah_id" gorm:"uniqueIndex"`
	Domain    string         `json:"domain" gorm:"type:varchar(255);uniqueIndex;not null"`
	IsActive  bool           `json:"is_active" gorm:"default:false;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Sekolah Sekolah `gorm:"foreignKey:SekolahID;constraint:OnDelete:CASCADE;" json:"sekolah"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE;"`
}

type Invoice struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	NoInvoice     string         `json:"no_invoice" gorm:"type:varchar(50);uniqueIndex;not null"`
	Tanggal       time.Time      `json:"tanggal" gorm:"type:date;not null"`
	StatusPayment string         `json:"status_payment" gorm:"type:varchar(20);default:'unpaid';not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
