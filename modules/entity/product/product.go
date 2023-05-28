package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_id             uint                `json:"ProductId"`
	Product_category_id    uint                `json:"product_category_id" form:"product_category_id"`
	Product_description_id uint                `json:"product_description_id" form:"product_description_id"`
	Name                   string              `json:"name" form:"name"`
	Stock                  int                 `json:"stock" form:"stock"`
	Price                  float64             `json:"price" form:"price"`
	Status                 string              `json:"status" form:"status"`
	Rating                 float64             `json:"rating" form:"rating"`
	Product_Category       Product_Category    `gorm:"foreignKey:Product_category_id"`
	Product_Description    Product_Description `gorm:"foreignKey:Product_description_id"`
}

type ProductRequest struct {
	Product_category_id uint     `json:"product_category_id" form:"product_category_id"`
	Name                string   `json:"name" form:"name"`
	Stock               int      `json:"stock" form:"stock"`
	Price               float64  `json:"price" form:"price"`
	Description         string   `json:"description" form:"description"`
	Status              string   `json:"status" form:"status"`
	Product_image_url   []string `json:"product_image_url" form:"product_image_url"`
}

type ProductResponse struct {
	Product_id        uint     `json:"product_id" form:"product_id"`
	Name              string   `json:"name" form:"name"`
	Category          string   `json:"category" form:"category"`
	Stock             int      `json:"stock" form:"stock"`
	Price             float64  `json:"price" form:"price"`
	Status            string   `json:"status" form:"status"`
	Rating            float64  `json:"rating" form:"rating"`
	Description       string   `json:"description" form:"description"`
	Product_image_url []string `json:"product_image_url" form:"product_image_url"`
}
