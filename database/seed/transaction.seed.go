package seed

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func CreateTransaction() []*et.Transaction {
	transaction := []*et.Transaction{
		{
			ExpeditionId:    1,
			VoucherId:       1,
			ShippingCost:    20000,
			PaymentMethodId: 3,
			AddressId:       2,
			Point:           5000,
			TotalPrice:      65000,
		},
		{
			ExpeditionId:    1,
			VoucherId:       1,
			ShippingCost:    20000,
			PaymentMethodId: 3,
			AddressId:       2,
			Point:           5000,
			TotalPrice:      65000,
		},
		{
			ExpeditionId:    1,
			VoucherId:       1,
			ShippingCost:    20000,
			PaymentMethodId: 3,
			AddressId:       2,
			Point:           5000,
			TotalPrice:      65000,
		},
	}

	return transaction
}
