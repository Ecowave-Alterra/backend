package transaction

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	"gorm.io/gorm"
)

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
	ExpeditionStatus   string `json:"ExpeditionStatus" form:"ExpeditionStatus"`
	CanceledReason     string
	ExpeditionRating   float32 `json:"ExpeditionRating" form:"ExpeditionRating"`
	Discount           float64 `json:"Discount" form:"Discount"`
	TotalPrice         float64
	TransactionDetails []TransactionDetail `json:"TransactionDetails" form:"TransactionDetails" gorm:"foreignKey:TransactionId"`
	Review             er.Review           `gorm:"foreignKey:TransactionId"`
}

type CanceledReason struct {
	CanceledReason string `json:"CanceledReason" form:"CanceledReason" validate:"required"`
}

type ShippingRequest struct {
	Weight float32 `json:"Weight" validate:"required"`
	CityId string  `json:"CityId" validate:"required"`
}
type ShippingResponse struct {
	Rajaongkir struct {
		Results []struct {
			Code  string `json:"Code"`
			Name  string `json:"Name"`
			Costs []struct {
				Service     string `json:"Service"`
				Description string `json:"Description"`
				Cost        []struct {
					Value uint   `json:"Value"`
					Etd   string `json:"Etd"`
				} `json:"Cost"`
			} `json:"Costs"`
		} `json:"Results"`
	} `json:"Rajaongkir"`
}
