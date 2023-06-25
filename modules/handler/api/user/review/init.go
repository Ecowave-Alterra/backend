package review

import (
	rc "github.com/berrylradianh/ecowave-go/modules/usecase/user/review"
)

type ReviewHandler struct {
	reviewUsecase rc.ReviewUsecase
}

func New(reviewUsecase rc.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{
		reviewUsecase,
	}
}
