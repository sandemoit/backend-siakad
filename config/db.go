package config

import (
	"fmt"
	"log"
	"os"
	"siakad/api/models"
	"siakad/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.LogError(err)
	}

	db.AutoMigrate(&models.User{}, &models.Santri{}, &models.Sekolah{}, &models.Guru{}, &models.Tenant{}, &models.Invoice{})

	DB = db

	postgresDB, err := db.DB()
	if err != nil {
		utils.LogError(err)
	}

	if err := postgresDB.Ping(); err != nil {
		utils.LogError(err)
	}

	utils.LogInfo("✅ Connected to PostgreSQL Server!")

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting DB instance:", err)
		return
	}
	sqlDB.Close()
}
