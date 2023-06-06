package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserDetail() []*ut.UserDetail {
	userDetail := []*ut.UserDetail{
		{
			FullName:        "user1 fullname",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave/img/users/profile/profile.png",
			EcoPoint:        0,
			UserId:          1,
		},
		{
			FullName:        "user2 fullname",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave/img/users/profile/profile.png",
			EcoPoint:        0,
			UserId:          2,
		},
	}

	return userDetail
}
