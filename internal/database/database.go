package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"healthcare-appointment-system/internal/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("healthcare.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.Patient{},
		&models.Appointment{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
}
