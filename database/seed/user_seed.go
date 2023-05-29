package seed

import (
	userEntity "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUser() []*userEntity.User {
	user := []*userEntity.User{
		{
			RoleId:      2,
			Email:       "user1@gmail.com",
			GoogleId:    "qwertyuiop",
			Username:    "user1",
			PhoneNumber: "123456789101",
			Password:    "passuser1",
		},
		{
			RoleId:      2,
			Email:       "user2@gmail.com",
			GoogleId:    "asdfghjkl",
			Username:    "user2",
			PhoneNumber: "123456789101",
			Password:    "passuser2",
		},
	}

	return user
}
