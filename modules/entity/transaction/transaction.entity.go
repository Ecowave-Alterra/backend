package transaction

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model

	UserId             uint `validate:"required"`
	VoucherId          uint `json:"VoucherId" form:"VoucherId"`
	AddressId          uint `json:"AddressId" form:"AddressId" validate:"required"`
	StatusTransaction  string
	ReceiptNumber      string
	TransactionId      string  `validate:"required"`
	TotalProductPrice  float64 `validate:"required"`
	TotalShippingPrice float64 `json:"TotalShippingPrice" form:"TotalShippingPrice" validate:"required"`
	Point              float64 `json:"Point" form:"Point"`
	PaymentMethod      string
	PaymentStatus      string
	ExpeditionName     string `json:"ExpeditionName" form:"ExpeditionName" validate:"required"`
	EstimationDay      string `json:"EstimationDay" form:"EstimationDay"`
	PaymentUrl         string `validate:"required"`
	CanceledReason     string
	ExpeditionRating   float32             `json:"ExpeditionRating" form:"ExpeditionRating"`
	Discount           float64             `json:"Discount" form:"Discount"`
	TotalPrice         float64             `validate:"required"`
	TransactionDetails []TransactionDetail `json:"TransactionDetails" form:"TransactionDetails" gorm:"foreignKey:TransactionId"`
}
