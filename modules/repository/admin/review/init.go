package review

import (
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	GetAllProducts(offset, pageSize int) ([]re.GetAllReviewResponse, int64, error)
	GetProductReviewById(productId string, offset, pageSize int) ([]re.ReviewResponse, int64, error)
	SearchProduct(search string, offset, pageSize int) ([]re.GetAllReviewResponse, int64, error)
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
