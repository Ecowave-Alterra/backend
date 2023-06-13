package product

import "gorm.io/gorm"

type Product struct {
	*gorm.Model       `json:"-"`
	ProductID         string          `json:"ProductId"`
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
	ProductID       string   `json:"ProductId"`
	Name            string   `json:"Name" form:"name"`
	Category        string   `json:"Category" form:"category"`
	Stock           uint     `json:"Stock" form:"stock"`
	Price           float64  `json:"Price" form:"price"`
	Status          string   `json:"Status" form:"status"`
	Rating          float64  `json:"Rating" form:"rating"`
	Description     string   `json:"Description" form:"description"`
	ProductImageUrl []string `json:"ProductImageUrl" form:"productImageUrl"`
}
