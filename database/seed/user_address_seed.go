package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserAddress() []*ut.UserAddress {
	userAddress := []*ut.UserAddress{
		{
			Recipient:    "ibu user1",
			PhoneNumber:  "085123456789",
			Address:      "bantul, jogja",
			Note:         "rumah cat krem",
			ProvinceId:   "11",
			ProvinceName: "Jawa Timur",
			CityId:       "247",
			CityName:     "Madiun",
			Mark:         "Rumah",
			IsPrimary:    true,
			UserId:       1,
		},
		{
			Recipient:    "satpam user1",
			PhoneNumber:  "085123456789",
			Address:      "sleman, jogja",
			Note:         "titip ke satpam aja",
			ProvinceId:   "11",
			ProvinceName: "Jawa Timur",
			CityId:       "251",
			CityName:     "Magetan",
			Mark:         "Kantor",
			IsPrimary:    false,
			UserId:       1,
		},
	}

	return userAddress
}
