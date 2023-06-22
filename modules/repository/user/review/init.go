package review

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"

	// et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	CountTransactionDetail(transactionId string) (int, error)
	GetIdTransaction(transactionId string) (int, error)
	GetProductId(transactionId string) ([]string, error)
	GetIdTransactionDetail(idTransaction int, productId string) (int, error)

	CreateRatingProduct(ratingProduct *er.RatingProduct) error
	UpdateExpeditionRating(ratingExpedition float32, transactionId string) error

	GetPoint(id int) (int, error)
	UpdatePoint(idUser int, point int) error
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
