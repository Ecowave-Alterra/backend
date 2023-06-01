package product

import "gorm.io/gorm"

type ProductImage struct {
	*gorm.Model
	ProductId       uint    `json:"productId" form:"productId"`
	ProductImageUrl string  `json:"productImageUrl" form:"productImageUrl"`
	Product         Product `gorm:"foreignKey:ProductId"`
}
