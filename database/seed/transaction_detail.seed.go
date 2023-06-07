package seed

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func CreateTransactionDetail() []*et.TransactionDetail {
	transactionDetail := []*et.TransactionDetail{
		{
			ProductId:     2,
			Qty:           3,
			SubTotalPrice: 20000,
		},
		{
			ProductId:     3,
			Qty:           1,
			SubTotalPrice: 30000,
		},
	}

	return transactionDetail
}
