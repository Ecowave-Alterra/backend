package review

import (
	"mime/multipart"

	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	ur "github.com/berrylradianh/ecowave-go/modules/repository/user/review"
)

type ReviewUsecase interface {
	CreateReview(review *er.Review, fileHeader, videoHeader *multipart.FileHeader) error
}

type reviewUsecase struct {
	reviewRepo ur.ReviewRepo
}

func New(reviewRepo ur.ReviewRepo) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo,
	}
}
