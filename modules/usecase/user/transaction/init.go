package transaction

import (
	em "github.com/berrylradianh/ecowave-go/modules/entity/midtrans"
	er "github.com/berrylradianh/ecowave-go/modules/entity/rajaongkir"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	rt "github.com/berrylradianh/ecowave-go/modules/repository/user/transaction"
)

type TransactionUsecase interface {
	CreateTransaction(transaction *et.Transaction) (string, string, error)
	GetPoint(id uint) (interface{}, error)
	GetVoucherUser(id uint, offset int, pageSize int) (interface{}, int64, error)
	MidtransNotifications(midtransRequest *em.MidtransRequest) error
	GetPaymentStatus(id string) (string, error)
	ShippingOptions(ship *er.RajaongkirRequest) (interface{}, error)
}

type transactionUsecase struct {
	transactionRepo rt.TransactionRepo
}

func New(adminRepo rt.TransactionRepo) *transactionUsecase {
	return &transactionUsecase{
		adminRepo,
	}
}
