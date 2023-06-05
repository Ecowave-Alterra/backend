package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (interface{}, error) {

	transactionDetail := transaction.TransactionDetails
	var productCost float64

	for _, cost := range transactionDetail {
		productCost += cost.SubTotalPrice
	}

	transaction.StatusTransactionId = 3
	transaction.ProductCost = productCost

	res, err := tu.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, err
	}

	return res, nil
}
