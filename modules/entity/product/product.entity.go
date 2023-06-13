package product

import "gorm.io/gorm"

type Product struct {
	*gorm.Model       `json:"-"`
	ProductID         string `json:"ProductId"`
	Name              string `validate:"required,max=10"`
	Stock             uint
	Price             float64
	Status            string
	Rating            float64
	Description       string
	ProductCategoryId uint            `json:"-"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryId" json:"-"`
}

type ProductRequest struct {
	ProductCategoryId uint     `json:"productCategoryId" form:"productCategoryId"`
	Name              string   `json:"name" form:"name"`
	Stock             uint     `json:"stock" form:"stock"`
	Price             float64  `json:"price" form:"price"`
	Description       string   `json:"description" form:"description"`
	Status            string   `json:"status" form:"status"`
	ProductImageUrl   []string `json:"productImageUrl" form:"productImageUrl"`
}

type ProductResponse struct {
	ProductID       string
	Name            string
	Category        string
	Stock           uint
	Price           float64
	Status          string
	Rating          float64
	Description     string
	ProductImageUrl []string
}
