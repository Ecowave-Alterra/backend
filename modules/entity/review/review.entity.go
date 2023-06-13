package review

import (
	"gorm.io/gorm"
)

type Review struct {
	*gorm.Model
	Rating       float64 `json:"rating" form:"rating"`
	Comment      string  `json:"comment" form:"comment"`
	CommentAdmin string  `json:"commentAdmin" form:"commentAdmin"`
	PhotoUrl     string  `json:"photoUrl" form:"photoUrl"`
	VideoUrl     string  `json:"videoUrl" form:"photoUrl"`
}

type GetAllReviewResponse struct {
	ProductID uint
	Name      string
	Category  string
	ReviewQty uint
}

type ReviewResponse struct {
	TransactionID uint
	Rating        float64
	Comment       string
	CommentAdmin  string
	PhotoUrl      string
	VideoUrl      string
}
