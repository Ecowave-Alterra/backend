package seed

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func CreateTransactionDetail() []*et.TransactionDetail {
	transaction := []*et.TransactionDetail{
		{
			TransactionId: 1,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductID:     1,
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 2,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductID:     2,
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 3,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductID:     2,
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 4,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductID:     3,
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
		{
			TransactionId: 5,
			ProductId:     "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductID:     3,
			ProductName:   "Product Name 1",
			Qty:           1,
			SubTotalPrice: 30000,
		},
	}
	return transaction
}
