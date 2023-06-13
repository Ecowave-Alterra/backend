package review

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type Review struct {
	*gorm.Model

	Rating            float64              `json:"Rating" form:"Rating"`
	Comment           string               `json:"Comment" form:"Comment"`
	CommentAdmin      string               `json:"CommentAdmin" form:"CommentAdmin"`
	PhotoUrl          string               `json:"PhotoUrl" form:"PhotoUrl"`
	VideoUrl          string               `json:"VideoUrl" form:"VideoUrl"`
	TransactionDetail et.TransactionDetail `gorm:"foreignKey:RatingProductId"`
}
