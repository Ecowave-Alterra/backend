package product

import "gorm.io/gorm"

type Product struct {
	*gorm.Model         `json:"-"`
	Name                string  `json:"name,omitempty" form:"name" validate:"required,max=10"`
	Stock               uint    `json:"stock,omitempty" form:"stock"`
	Price               float64 `json:"price,omitempty" form:"price"`
	Status              string  `json:"status" form:"status"`
	Rating              float64 `json:"rating" form:"rating"`
	Description         string  `json:"description,omitempty" form:"description"`
	Product_category_id uint    `json:"product_category_id,omitempty" form:"product_category_id"`
	// Product_description_id uint               `json:"product_description_id,omitempty" form:"product_description_id"`
	Product_Category ProductCategory `gorm:"foreignKey:Product_category_id"`
	// Product_Description    ProductDescription `gorm:"foreignKey:Product_description_id"`
}

type ProductRequest struct {
	Product_category_id uint     `json:"product_category_id" form:"product_category_id"`
	Name                string   `json:"name" form:"name"`
	Stock               uint     `json:"stock" form:"stock"`
	Price               float64  `json:"price" form:"price"`
	Description         string   `json:"description" form:"description"`
	Status              string   `json:"status" form:"status"`
	Product_image_url   []string `json:"product_image_url" form:"product_image_url"`
}

type ProductResponse struct {
	Product_id        uint     `json:"product_id"`
	Name              string   `json:"name,omitempty" form:"name"`
	Category          string   `json:"category,omitempty" form:"category"`
	Stock             uint     `json:"stock,omitempty" form:"stock"`
	Price             float64  `json:"price,omitempty" form:"price"`
	Status            string   `json:"status,omitempty" form:"status"`
	Rating            float64  `json:"rating,omitempty" form:"rating"`
	Description       string   `json:"description,omitempty" form:"description"`
	Product_image_url []string `json:"product_image_url,omitempty" form:"product_image_url"`
}
