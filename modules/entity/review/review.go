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

type Review struct {
	*gorm.Model

	RatingService float64        `json:"RatingService" form:"RatingService"`
	TransactionId uint           `json:"TransactionId" form:"TransactionId"`
	ReviewDetails []ReviewDetail `gorm:"foreignKey:ReviewId"`
}

type ReviewDetail struct {
	*gorm.Model

	RatingProduct float64 `json:"RatingProduct" form:"RatingProduct"`
	Comment       string  `json:"Comment" form:"Comment"`
	CommentAdmin  string  `json:"CommentAdmin" form:"CommentAdmin"`
	PhotoUrl      string  `json:"PhotoUrl" form:"PhotoUrl"`
	VideoUrl      string  `json:"VideoUrl" form:"VideoUrl"`
	ReviewId      uint    `json:"ReviewId" form:"ReviewId"`
}
