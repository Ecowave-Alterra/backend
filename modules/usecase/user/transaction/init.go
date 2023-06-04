package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	rt "github.com/berrylradianh/ecowave-go/modules/repository/user/transaction"
)

type TransactionUsecase interface {
	CreateTransaction(transaction *et.Transaction) (interface{}, error)
}

type transactionUsecase struct {
	transactionRepo rt.TransactionRepo
}

func New(adminRepo rt.TransactionRepo) *transactionUsecase {
	return &transactionUsecase{
		adminRepo,
	}
}
