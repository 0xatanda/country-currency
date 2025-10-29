package database

import (
	"log"

	"github.com/0xatanda/country-currency/internal/config"
	"github.com/0xatanda/country-currency/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectProgres() (*gorm.DB, error) {
	cfg := config.LoadDBConfig()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Country{}); err != nil {
		log.Fatalf("Migration failed %v", err)
	}

	return db, nil
}
