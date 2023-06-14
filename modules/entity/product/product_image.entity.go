package product

import "gorm.io/gorm"

type ProductImage struct {
	*gorm.Model     `json:"-"`
	ProductId       uint    `json:"-"`
	ProductImageUrl string  `json:"ProductImageUrl" form:"ProductImageUrl"`
	Product         Product `json:"-"`
}
