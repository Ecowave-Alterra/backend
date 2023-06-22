package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type OrderRepo interface {
	GetAllOrder(transaction *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error)
}

type orderRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db,
	}
}
