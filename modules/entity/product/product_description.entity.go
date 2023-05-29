package product

import "gorm.io/gorm"

type ProductDescription struct {
	*gorm.Model `json:"-"`

	Description string `json:"description" form:"description" validate:"required"`
	// Products    []ProductResponse `gorm:"foreignKey:Product_description_id"`
}

type ProductDescriptionResponse struct {
	Description string `json:"description" form:"description"`
	// Products    []Product `gorm:"foreignKey:Product_description_id" json:"-"`
}
