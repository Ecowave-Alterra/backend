package product

import "gorm.io/gorm"

type ProductCategory struct {
	*gorm.Model `json:"-"`

	ID       uint      `json:"Id,omitempty" gorm:"primary_key"`
	Category string    `json:"Category" form:"category" validate:"required"`
	Products []Product `gorm:"foreignKey:ProductCategoryId"`
}

type ProductCategoryResponse struct {
	Category string    `json:"category" form:"category"`
	Products []Product `gorm:"foreignKey:ProductCategoryId" json:"-"`
}
