package transaction

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (tr *transactionRepo) CreateTransaction(transaction *et.Transaction) (interface{}, error) {

	err := tr.db.Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
