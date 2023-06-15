package review

import (
	"mime/multipart"

	ur "github.com/berrylradianh/ecowave-go/modules/repository/user/review"
)

type ReviewUsecase interface {
	CountTransactionDetail(transactionId string) (int, error)
	GetIdTransactionDetail(transactionId string) ([]int, error)
	CreateRatingProduct(rating float64, comment string, fileHeader, videoHeader *multipart.FileHeader, transactionDetailId int) error
	UpdateExpeditionRating(ratingExpedition float32, transactionId string) error
	UpdatePoint(idUser int) error
}

type reviewUsecase struct {
	reviewRepo ur.ReviewRepo
}

func New(reviewRepo ur.ReviewRepo) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo,
	}
}
