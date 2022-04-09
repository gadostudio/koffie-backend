package products

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id           string `json:"item_id"`
	Name         string `json:"name"`
	Desc         string `json:"description"`
	Price        uint   `json:"price"`
	Image        string `json:"image_url"`
	Undiscounted uint   `json:"undiscounted_price"`
}
