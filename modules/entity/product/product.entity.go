package product

import "gorm.io/gorm"

type Product struct {
	*gorm.Model `json:"-"`

	Name      string  `json:"name,omitempty" form:"name"`
	Price     float64 `json:"price,omitempty" form:"price"`
	Stock     int     `json:"stock,omitempty" form:"stock"`
	Image_url string  `json:"image_url,omitempty" form:"image_url"`

	Product_category_id uint                    `json:"product_category_id,omitempty" form:"product_category_id"`
	Product_category    ProductCategoryResponse `gorm:"foreignKey:Product_category_id"`
}

type ProductResponse struct {
	*gorm.Model         `json:"-"`
	Name                string  `json:"name,omitempty" form:"name"`
	Price               float64 `json:"price,omitempty" form:"price"`
	Stock               int     `json:"stock,omitempty" form:"stock"`
	Image_url           string  `json:"image_url,omitempty" form:"image_url"`
	Product_category_id uint    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}
