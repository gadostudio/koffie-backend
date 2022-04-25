package db

import (
	"github.com/shaderboi/koffie-backend/api/midtrans"
	"github.com/shaderboi/koffie-backend/api/products"
	"github.com/shaderboi/koffie-backend/api/stores"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connection() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	db.AutoMigrate(&products.Product{})
	db.AutoMigrate(&midtrans.Payment{})
	db.AutoMigrate(&products.Category{})
	db.AutoMigrate(&stores.Store{})

	return db, err

}
