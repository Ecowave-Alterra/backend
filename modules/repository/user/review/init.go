package review

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	// et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	CreateReview(review *er.Review) error
	CreateReviewDetail(reviewDetail *er.ReviewDetail) error
	GetIdReview(idTransaction int) (int, error)
	GetIdTransaction(transactionId string) (int, error)
	CountTransactionDetail(transactionId string) (int, error)
	GetPoint(id int) (int, error)
	UpdatePoint(id int, point int) error
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
