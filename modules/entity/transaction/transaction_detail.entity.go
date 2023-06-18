package transaction

import "gorm.io/gorm"

type TransactionDetail struct {
	*gorm.Model

	TransactionId   uint
	ProductID       uint
	ProductId       string  `json:"ProductId" form:"ProductId" validate:"required"`
	ProductName     string  `json:"ProductName" form:"ProductName" validate:"required"`
	RatingProductId uint    `json:"RatingProductId" form:"RatingProductId"`
	Qty             uint    `json:"Qty" form:"Qty" validate:"required"`
	SubTotalPrice   float64 `json:"SubTotalPrice" form:"SubTotalPrice" validate:"required"`
}
