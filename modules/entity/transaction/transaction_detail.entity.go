package transaction

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"

	"gorm.io/gorm"
)

type TransactionDetail struct {
	*gorm.Model

	TransactionId uint
	ProductId     uint             `json:"ProductId" form:"ProductId" validate:"required"`
	Qty           uint             `json:"Qty" form:"Qty" validate:"required"`
	SubTotalPrice float64          `json:"SubTotalPrice" form:"SubTotalPrice" validate:"required"`
	RatingProduct er.RatingProduct `gorm:"foreignKey:TransactionDetailId"`
}
