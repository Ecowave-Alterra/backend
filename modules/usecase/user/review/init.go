package review

import (
	"mime/multipart"

	// et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ur "github.com/berrylradianh/ecowave-go/modules/repository/user/review"
)

type ReviewUsecase interface {
	CreateReview(ratingService float64, transactionId string) error
	CreateReviewDetail(ratingProduct float64, comment string, fileHeader, videoHeader *multipart.FileHeader, transactionId string) error
	CountTransactionDetail(transactionId string) (int, error)
	UpdatePoint(id int) error
}

type reviewUsecase struct {
	reviewRepo ur.ReviewRepo
}

func New(reviewRepo ur.ReviewRepo) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo,
	}
}
