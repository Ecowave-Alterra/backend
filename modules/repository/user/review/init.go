package review

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	CreateReview(review *er.Review) error
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
