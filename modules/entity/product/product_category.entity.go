package product

import "gorm.io/gorm"

type ProductCategory struct {
	*gorm.Model `json:"-"`

	Name     string            `json:"name" form:"name" validate:"required"`
	Products []ProductResponse `gorm:"foreignKey:Product_category_id"`
}

type ProductCategoryResponse struct {
	*gorm.Model `json:"-"`
	Name        string    `json:"name" form:"name"`
	Products    []Product `gorm:"foreignKey:Product_category_id" json:"-"`
}

func (ProductCategoryResponse) TableName() string {
	return "product_categories"
}
