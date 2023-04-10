package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/moaabb/product-api/pkg/models"
)

type Handler struct {
	DB *gorm.DB
}

func Init(dsn string) Handler {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Product{})

	return Handler{db}
}
