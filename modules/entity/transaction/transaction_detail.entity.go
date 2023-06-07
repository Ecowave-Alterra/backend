package transaction

import "gorm.io/gorm"

type TransactionDetail struct {
	*gorm.Model

	TransactionId   uint
	ProductId       uint    `json:"ProductId" form:"ProductId" validate:"required"`
	RatingProductId uint    `json:"RatingProductId" form:"RatingProductId"`
	Qty             uint    `json:"Qty" form:"Qty" validate:"required"`
	SubTotalPrice   float64 `json:"SubTotalPrice" form:"SubTotalPrice" validate:"required"`
}
