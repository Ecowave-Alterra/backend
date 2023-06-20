package seed

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func CreateTransactionDetail() []*et.TransactionDetail {
	transaction := []*et.TransactionDetail{
		{
			TransactionId: 1,
			ProductId:     "859679ad-888b-41ac-a663-4a619acc4ae3",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 2,
			ProductId:     "859679ad-888b-41ac-a663-4a619acc4ae3",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 3,
			ProductId:     "859679ad-888b-41ac-a663-4a619acc4ae3",
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 4,
			ProductId:     "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductName:   "Product Name 2",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 5,
			ProductId:     "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductName:   "Product Name 2",
			Qty:           1,
			SubTotalPrice: 30000,
		},
	}
	return transaction
}
