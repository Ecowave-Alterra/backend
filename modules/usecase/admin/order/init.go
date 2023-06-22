package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	or "github.com/berrylradianh/ecowave-go/modules/repository/admin/order"
)

type OrderUseCase interface {
	GetAllOrder(transaction *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error)
}

type orderUseCase struct {
	orderRepo or.OrderRepo
}

func New(orderRepo or.OrderRepo) *orderUseCase {
	return &orderUseCase{
		orderRepo,
	}
}
