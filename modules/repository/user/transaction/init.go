package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transaction *et.Transaction) (interface{}, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TransactionRepo {
	return &transactionRepo{
		db,
	}
}
