package seed

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func CreateTransactionDetail() []*et.TransactionDetail {
	transactionDetail := []*et.TransactionDetail{
		{
			TransactionId:   1,
			ProductId:       1,
			RatingProductId: 1,
			Qty:             3,
			SubTotalPrice:   20000,
		},
		{
			TransactionId:   1,
			ProductId:       2,
			RatingProductId: 2,
			Qty:             1,
			SubTotalPrice:   30000,
		},
		{
			TransactionId:   2,
			ProductId:       1,
			RatingProductId: 3,
			Qty:             3,
			SubTotalPrice:   20000,
		},
		{
			TransactionId:   3,
			ProductId:       1,
			RatingProductId: 4,
			Qty:             3,
			SubTotalPrice:   20000,
		},
	}

	return transactionDetail
}
