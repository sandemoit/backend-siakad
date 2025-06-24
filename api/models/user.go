package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100"`
	Email    string `gorm:"unique"`
	Password string
	Role     string `gorm:"size:50"` // admin, santri, ustadz, wali
}
