package product

import "gorm.io/gorm"

type Product struct {
	*gorm.Model       `json:"-"`
	Name              string          `json:"name" form:"name" validate:"required,max=10"`
	Stock             uint            `json:"stock" form:"stock"`
	Price             float64         `json:"price" form:"price"`
	Status            string          `json:"status" form:"status"`
	Rating            float64         `json:"rating" form:"rating"`
	Description       string          `json:"description" form:"description"`
	ProductCategoryId uint            `json:"productCategoryId" form:"productCategoryId"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryId"`
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
	ProductId       uint     `json:"product_id"`
	Name            string   `json:"name" form:"name"`
	Category        string   `json:"category" form:"category"`
	Stock           uint     `json:"stock" form:"stock"`
	Price           float64  `json:"price" form:"price"`
	Status          string   `json:"status" form:"status"`
	Rating          float64  `json:"rating" form:"rating"`
	Description     string   `json:"description" form:"description"`
	ProductImageUrl []string `json:"productImageUrl" form:"productImageUrl"`
}
