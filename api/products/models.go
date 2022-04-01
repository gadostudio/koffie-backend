package products

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Name  string
	Price uint
	Stock uint
}
