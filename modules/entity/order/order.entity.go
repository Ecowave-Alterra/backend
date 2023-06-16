package order

import (
	"time"

	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

type Order struct {
	TransactionId      string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	AddressId          uint
	StatusTransaction  string
	ReceiptNumber      string
	TotalProductPrice  float64
	TotalShippingPrice float64
	Point              float64
	PaymentMethod      string
	PaymentStatus      string
	ExpeditionName     string
	ExpeditionStatus   string
	VoucherId          uint
	Discount           float64
	TotalPrice         float64
	CanceledReason     string
	OrderDetail        []OrderDetail
	Address            eu.UserAddress
}

type OrderDetail struct {
	// ProductId       string
	ProductId       uint
	Qty             uint
	SubTotalPrice   float64
	ProductName     string
	ProductImageUrl string
}
