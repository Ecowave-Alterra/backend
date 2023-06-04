package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (interface{}, error) {

	transaction.StatusTransactionId = 3

	res, err := tu.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, err
	}

	return res, nil
}
