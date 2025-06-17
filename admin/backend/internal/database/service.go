package database

import (
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
	"fmt"
	"admin/internal/database/models"
)

var DB *gorm.DB


func InitDatabase() error {
	var err error

	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&models.UserActivity{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

