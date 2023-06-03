package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserDetail() []*ut.UserDetail {
	userDetail := []*ut.UserDetail{
		{
			FullName:        "user1 fullname",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave_storage/img/swis.jpg",
			EcoPoint:        0,
			UserId:          1,
		},
		{
			FullName:        "user2 fullname",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave_storage/img/swis.jpg",
			EcoPoint:        0,
			UserId:          2,
		},
	}

	return userDetail
}
