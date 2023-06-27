package review

import (
	rc "github.com/berrylradianh/ecowave-go/modules/usecase/admin/review"
)

type ReviewHandler struct {
	reviewUsecase rc.ReviewUseCase
}

func New(reviewUsecase rc.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{
		reviewUsecase,
	}
}
