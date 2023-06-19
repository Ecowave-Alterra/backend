package transaction

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model

	UserId             uint
	VoucherId          uint `json:"VoucherId" form:"VoucherId"`
	AddressId          uint `json:"AddressId" form:"AddressId"`
	StatusTransaction  string
	ReceiptNumber      string
	TransactionId      string `json:"TransactionId" form:"TransactionId"`
	TotalProductPrice  float64
	TotalShippingPrice float64 `json:"TotalShippingPrice" form:"TotalShippingPrice" validate:"required"`
	Point              float64 `json:"Point" form:"Point"`
	PaymentMethod      string  `json:"PaymentMethod" form:"PaymentMethod" validate:"required"`
	PaymentStatus      string
	ExpeditionName     string `json:"ExpeditionName" form:"ExpeditionName" validate:"required"`
	EstimationDay      string `json:"EstimationDay" form:"EstimationDay"`
	PaymentUrl         string `json:"PaymentUrl" form:"PaymentUrl"`
	CanceledReason     string
	ExpeditionRating   float32 `json:"ExpeditionRating" form:"ExpeditionRating"`
	Discount           float64 `json:"Discount" form:"Discount"`
	TotalPrice         float64
	TransactionDetails []TransactionDetail `json:"TransactionDetails" form:"TransactionDetails" gorm:"foreignKey:TransactionId"`
}
