package product

import "gorm.io/gorm"

type ProductCategory struct {
	*gorm.Model `json:"-"`

	Category string    `json:"category" form:"category" validate:"required"`
	Products []Product `gorm:"foreignKey:Product_category_id"`
}

type ProductCategoryResponse struct {
	Category string    `json:"category" form:"category"`
	Products []Product `gorm:"foreignKey:Product_category_id" json:"-"`
}
