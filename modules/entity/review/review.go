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

type GetAllReviewResponse struct {
	ProductID string
	Name      string
	Category  string
	ReviewQty uint
}

type ReviewResponse struct {
	TransactionID    string
	Name             string
	ProfilePhoto     string
	ProductName      string
	ProductCategory  string
	CommentUser      string
	CommentAdmin     string
	PhotoUrl         string
	VideoUrl         string
	AvgRating        float64
	ExpeditionRating float32
	ProductRating    float64
}
