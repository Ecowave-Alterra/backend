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
	PaymentMethod      string  `json:"PaymentMethodId" form:"PaymentMethodId" validate:"required"`
	PaymentStatus      string  `json:"PaymentStatus" form:"PaymentStatus" validate:"required"`
	ExpeditionName     string  `json:"ExpeditionName" form:"ExpeditionName" validate:"required"`
	ExpeditionStatus   string  `json:"ExpeditionStatus" form:"ExpeditionStatus"`
	CanceledReason     string
	ExpeditionRating   float32 `json:"ExpeditionRating" form:"ExpeditionRating"`
	Discount           float64 `json:"Discount" form:"Discount"`
	TotalPrice         float64
	TransactionDetails []TransactionDetail `json:"TransactionDetails" form:"TransactionDetails" gorm:"foreignKey:TransactionId"`
}
type CanceledReason struct {
	CanceledReason string `json:"CanceledReason" form:"CanceledReason" validate:"required"`
}
