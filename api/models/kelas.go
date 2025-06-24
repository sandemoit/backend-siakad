package models

type Kelas struct {
	ID      uint   `gorm:"primaryKey"`
	Nama    string `gorm:"size:100"`
	Jenjang string `gorm:"size:50"`
}
