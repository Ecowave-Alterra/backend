package review

import (
	"gorm.io/gorm"
)

type RatingProduct struct {
	*gorm.Model

	Rating              float64 `json:"Rating" form:"Rating"`
	Comment             string  `json:"Comment" form:"Comment"`
	CommentAdmin        string  `json:"CommentAdmin" form:"CommentAdmin"`
	PhotoUrl            string  `json:"PhotoUrl" form:"PhotoUrl"`
	VideoUrl            string  `json:"VideoUrl" form:"VideoUrl"`
	TransactionDetailId uint    `json:"TransactionDetailId" form:"TransactionDetailId"`
}
