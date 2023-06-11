package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	rt "github.com/berrylradianh/ecowave-go/modules/repository/user/transaction"
)

type TransactionUsecase interface {
	CreateTransaction(transaction *et.Transaction) error
	GetPoint(id uint) (uint, error)
	ClaimVoucher(idUser uint, idVoucher uint, shipCost float64, productCost float64) (float64, error)
	GetVoucherUser(id uint, offset int, pageSize int) (interface{}, int64, error)
	DetailVoucher(id uint) (interface{}, error)
}

type transactionUsecase struct {
	transactionRepo rt.TransactionRepo
}

func New(adminRepo rt.TransactionRepo) *transactionUsecase {
	return &transactionUsecase{
		adminRepo,
	}
}
