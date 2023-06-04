package transaction

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model

	UserId              uint
	PaymentMethodId     uint `json:"PaymentMethodId" form:"PaymentMethodId" validate:"required"`
	ExpeditionId        uint `json:"ExpeditionId" form:"ExpeditionId" validate:"required"`
	VoucherId           uint `json:"VoucherId" form:"VoucherId"`
	AddressId           uint `json:"AddressId" form:"AddressId"`
	StatusTransactionId uint
	RatingExpeditionId  uint                `json:"RatingExpeditionId" form:"RatingExpeditionId"`
	ShippingCost        float64             `json:"ShippingCost" form:"ShippingCost" validate:"required"`
	Point               float64             `json:"Point" form:"Point"`
	TotalPrice          float64             `json:"TotalPrice" form:"TotalPrice" validate:"required"`
	TransactionDetails  []TransactionDetail `json:"TransactionDetails" form:"TransactionDetails" gorm:"foreignKey:TransactionId"`
}
