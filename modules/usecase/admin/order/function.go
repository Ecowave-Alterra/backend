package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (oc *orderUseCase) GetAllOrder(transaction *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error) {
	return oc.orderRepo.GetAllOrder(transaction, offset, pageSize)
}

func (oc *orderUseCase) GetOrderByID(transactionId string, transaction *te.TransactionDetailResponse) (te.TransactionDetailResponse, error) {
	return oc.orderRepo.GetOrderByID(transactionId, transaction)
}

func (oc *orderUseCase) GetOrderProducts(transactionId string, products *[]te.TransactionProductDetailResponse) ([]te.TransactionProductDetailResponse, error) {
	return oc.orderRepo.GetOrderProducts(transactionId, products)
}

func (oc *orderUseCase) SearchOrder(search, filter string, offset, pageSize int) (*[]te.TransactionResponse, int64, error) {
	return oc.orderRepo.SearchOrder(search, filter, offset, pageSize)
}

func (oc *orderUseCase) UpdateReceiptNumber(transactionId string, receiptNumber string) error {
	return oc.orderRepo.UpdateReceiptNumber(transactionId, receiptNumber)
}
