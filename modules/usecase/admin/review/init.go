package review

import (
	rr "github.com/berrylradianh/ecowave-go/modules/repository/admin/review"
)

type ReviewUseCase interface {
}

type reviewUsecase struct {
	reviewRepo rr.ReviewRepo
}

func New(reviewRepo rr.ReviewRepo) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo,
	}
}
