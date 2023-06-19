package seed

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUserDetail() []*ut.UserDetail {
	userDetail := []*ut.UserDetail{
		{
			Name:         "Administrator",
			Point:        0,
			Phone:        "0897162534215",
			ProfilePhoto: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:       1,
		},
		{
			Name:         "User 1",
			Point:        0,
			Phone:        "08917283129283",
			ProfilePhoto: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:       2,
		},
		{
			Name:         "User 2",
			Point:        0,
			Phone:        "0851728392716",
			ProfilePhoto: "https://storage.googleapis.com/ecowave/img/users/profile/profile.png",
			UserId:       3,
		},
	}

	return userDetail
}
