package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserAddress() []*ut.UserAddress {
	userAddress := []*ut.UserAddress{
		{
			Recipient:    "Administrator",
			Phone:        "08917283129283",
			ProvinceId:   "11",
			ProvinceName: "Jawa Timur",
			CityId:       "255",
			CityName:     "Malang",
			Address:      "Sukun",
			Note:         "Rumah Kuning",
			Mark:         "Rumah",
			IsPrimary:    true,
			UserId:       1,
		},
		{
			Recipient:    "Ibu user1",
			Phone:        "085123456789",
			ProvinceId:   "11",
			ProvinceName: "Jawa Timur",
			CityId:       "247",
			CityName:     "Madiun",
			Address:      "Balerejo",
			Note:         "Rumah cat krem",
			Mark:         "Rumah",
			IsPrimary:    true,
			UserId:       2,
		},
		{
			Recipient:    "Satpam user1",
			Phone:        "085123456789",
			ProvinceId:   "5",
			ProvinceName: "DI Yogyakarta",
			CityId:       "39",
			CityName:     "Sleman",
			Address:      "Gentan",
			Note:         "Titip ke mas kos",
			Mark:         "Rumah",
			IsPrimary:    false,
			UserId:       2,
		},
	}

	return userAddress
}
