package review

import (
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	rr "github.com/berrylradianh/ecowave-go/modules/repository/admin/review"
)

type ReviewUseCase interface {
	GetAllProducts(offset, pageSize int) ([]re.GetAllReviewResponse, int64, error)
	GetProductReviewById(productId string, offset, pageSize int) ([]re.ReviewResponse, int64, error)
	SearchProduct(search string, offset, pageSize int) ([]re.GetAllReviewResponse, int64, error)
}

type reviewUsecase struct {
	reviewRepo rr.ReviewRepo
}

func New(reviewRepo rr.ReviewRepo) *reviewUsecase {
	return &reviewUsecase{
		reviewRepo,
	}
}
