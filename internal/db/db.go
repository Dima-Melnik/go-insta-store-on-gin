package db

import (
	"fmt"

	"github.com/Dima-Melnik/go-insta-store-on-gin/config"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := config.Cfg.InitDatabaseDsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Attribute{}); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	DB = db
	return nil
}

func CloseDB() error {
	postDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed init database: %v", err)
	}

	if err := postDB.Close(); err != nil {
		return fmt.Errorf("failed to close db connect")
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
