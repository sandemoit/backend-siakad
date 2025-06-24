package models

type Santri struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100"`
	NISN         string `gorm:"unique"`
	KelasID      uint
	Alamat       string
	TempatLahir  string
	TanggalLahir string
	Foto         string
}
