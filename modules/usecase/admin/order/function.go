package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (oc *orderUseCase) GetAllOrder(transaction *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error) {
	return oc.orderRepo.GetAllOrder(transaction, offset, pageSize)
}
