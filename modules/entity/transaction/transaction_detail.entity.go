package transaction

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	*gorm.Model

	TransactionId   uint
	ProductId       uint       `json:"ProductId" form:"ProductId" validate:"required"`
	RatingProductId uint       `json:"RatingProductId" form:"RatingProductId"`
	Qty             uint       `json:"Qty" form:"Qty" validate:"required"`
	SubTotalPrice   float64    `json:"SubTotalPrice" form:"SubTotalPrice" validate:"required"`
	Product         pe.Product `gorm:"foreignKey:ProductId"`
	Review          re.Review  `gorm:"foreignKey:RatingProductId"`
}
