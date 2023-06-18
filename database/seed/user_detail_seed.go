package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserDetail() []*ut.UserDetail {
	userDetail := []*ut.UserDetail{
		{
			Name:            "user1 fullname",
			Point:           0,
			Phone:           "08998754321",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave/img/users/profile/profile.png",
			UserId:          1,
		},
		{
			Name:            "user2 fullname",
			Point:           0,
			Phone:           "08998754321",
			ProfilePhotoUrl: "https://storage.cloud.google.com/ecowave/img/users/profile/profile.png",
			UserId:          2,
		},
	}

	return userDetail
}
