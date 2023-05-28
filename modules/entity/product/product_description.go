package product

import "gorm.io/gorm"

type Product_Description struct {
	gorm.Model
	Description string `json:"description" form:"description"`
}
