package product

import "gorm.io/gorm"

type Product_Image struct {
	gorm.Model
	Product_id        uint    `json:"product_id" form:"product_id"`
	Product_image_url string  `json:"product_image_url" form:"product_image_url"`
	Product           Product `gorm:"foreignKey:Product_id"`
}
