package db

import (
	"github.com/shaderboi/koffie-backend/api/midtrans"
	"github.com/shaderboi/koffie-backend/api/products"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connection() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	db.AutoMigrate(&products.Product{})
	db.AutoMigrate(&midtrans.Payment{})

	return db, err

}
