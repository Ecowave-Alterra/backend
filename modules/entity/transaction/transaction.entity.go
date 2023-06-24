package transaction

import (
	"time"

	"gorm.io/gorm"
)

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

type TransactionResponse struct {
	ReceiptNumber     string
	TransactionId     string
	Name              string
	Unit              uint
	TotalPrice        float64
	OrderDate         time.Time
	StatusTransaction string
}

type TransactionDetailResponse struct {
	Address            string
	Voucher            string
	Name               string
	PhoneNumber        string
	ReceiptNumber      string
	TotalProductPrice  float64
	TotalShippingPrice float64
	TotalPrice         float64
	Point              float64
	PaymentMethod      string
	PaymentStatus      string
	CanceledReason     string
	ExpeditionRating   float32
	StatusTransaction  string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type TransactionProductDetailResponse struct {
	ProductName     string
	ProductImageUrl string
	Qty             string
}
