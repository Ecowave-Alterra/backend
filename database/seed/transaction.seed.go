package seed

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func CreateTransaction() []*et.Transaction {
	transaction := []*et.Transaction{
		{
			UserId:           2,
			ExpeditionId:     1,
			VoucherId:        1,
			ShippingCost:     20000,
			PaymentMethodId:  3,
			AddressId:        2,
			ExpeditionRating: 4.5,
			Point:            5000,
			TotalPrice:       65000,
		},
		{
			UserId:           2,
			ExpeditionId:     1,
			VoucherId:        1,
			ShippingCost:     20000,
			PaymentMethodId:  3,
			AddressId:        2,
			ExpeditionRating: 5,
			Point:            5000,
			TotalPrice:       65000,
		},
		{
			UserId:           1,
			ExpeditionId:     1,
			VoucherId:        1,
			ShippingCost:     20000,
			PaymentMethodId:  3,
			AddressId:        2,
			ExpeditionRating: 3.5,
			Point:            5000,
			TotalPrice:       65000,
		},
	}

	return transaction
}
