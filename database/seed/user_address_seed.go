package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserAddress() []*ut.UserAddress {
	userAddress := []*ut.UserAddress{
		{
			Recipient:   "ibu user1",
			PhoneNumber: "085123456789",
			Address:     "bantul, jogja",
			Note:        "rumah cat krem",
			Mark:        "Rumah",
			IsPrimary:   true,
			UserId:      1,
		},
		{
			Recipient:   "satpam user1",
			PhoneNumber: "085123456789",
			Address:     "sleman, jogja",
			Note:        "titip ke satpam aja",
			Mark:        "Kantor",
			IsPrimary:   false,
			UserId:      1,
		},
	}

	return userAddress
}
