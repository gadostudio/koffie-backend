package products

import (
	"time"
)

type Product struct {
	Id         uint       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name       string     `json:"name"`
	Desc       string     `json:"description"`
	Price      uint       `json:"price"`
	Image      string     `json:"image_url"`
	CategoryId uint       `json:"category_id" gorm:"foreignKey"`
	Discount   *uint      `json:"discount"`
	CreatedAt  *time.Time `json:"created_at"`
}

type Category struct {
	Id           uint       `gorm:"primaryKey;AUTO_INCREMENT" json:"category_id"`
	CategoryName string     `json:"category_name"`
	CreatedAt    *time.Time `json:"created_at"`
}

type CategorizedProduct struct {
	CategoryName string
	Category     Category  `gorm:"references:CategoryName" json:"category"`
	Products     []Product `json:"products"`
}

type Products struct {
	ProductList  []Product
	CategoryList []Category
}

func (p Products) Categorize() []CategorizedProduct {

	var cat_products []CategorizedProduct

	var products []Product

	for _, cat := range p.CategoryList {
		category_name := cat.CategoryName

		for _, prod := range p.ProductList {
			if cat.Id == prod.CategoryId {
				products = append(products, prod)
			}
		}

		product := CategorizedProduct{
			CategoryName: category_name,
			Products:     products,
		}

		cat_products = append(cat_products, product)
	}

	return cat_products
}
