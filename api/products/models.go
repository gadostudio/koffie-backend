package products

import (
	"time"
)

type Product struct {
	Id        string     `gorm:"primaryKey" json:"item_id"`
	Name      string     `json:"name"`
	Desc      string     `json:"description"`
	Price     int        `json:"price"`
	Image     string     `json:"image_url"`
	Discount  *int       `json:"discount"`
	CreatedAt *time.Time `json:"created_at"`
}
