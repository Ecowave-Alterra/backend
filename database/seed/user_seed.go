package seed

import (
	userEntity "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUser() []*userEntity.UserRequest {
	user := []*userEntity.UserRequest{
		{
			Name:        "user1",
			Email:       "user1@gmail.com",
			Username:    "user1",
			PhoneNumber: "123456789101",
			Password:    "passuser1",
		},
		{
			Name:        "user2",
			Email:       "user2@gmail.com",
			Username:    "user2",
			PhoneNumber: "123456789101",
			Password:    "passuser2",
		},
	}

	return user
}
