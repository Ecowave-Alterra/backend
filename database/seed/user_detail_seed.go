package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserDetail() []*ut.UserDetail {
	userDetail := []*ut.UserDetail{
		{
			Name:            "user1 fullname",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave/img/users/profile/profile.png",
			Point:           0,
			UserId:          1,
		},
		{
			Name:            "user2 fullname",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave/img/users/profile/profile.png",
			Point:           0,
			UserId:          2,
		},
	}

	return userDetail
}
