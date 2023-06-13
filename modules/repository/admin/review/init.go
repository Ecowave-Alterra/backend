package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	// re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	GetAllProducts(products *[]pe.Product) ([]pe.Product, error)
	GetAllTransactionDetails(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error)
}

type reviewRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db,
	}
}
