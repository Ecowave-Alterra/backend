package product

import "gorm.io/gorm"

type Product_Category struct {
	gorm.Model
	Category string `json:"category" form:"category"`
}
