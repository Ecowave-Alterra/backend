package product

import "gorm.io/gorm"

type ProductImage struct {
	*gorm.Model     `json:"-"`
	ProductId       string  `gorm:"size:255" json:"-"`
	ProductImageUrl string  `json:"ProductImageUrl" form:"ProductImageUrl"`
	Product         Product `json:"-"`
}
