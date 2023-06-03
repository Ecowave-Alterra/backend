package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUser() []*ut.User {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")

	user := []*ut.User{
		{
			RoleId:      2,
			Email:       "user1@gmail.com",
			GoogleId:    "qwertyuiop",
			Username:    "user1",
			PhoneNumber: "085123456789",
			Password:    string(hashPasswordUser1),
		},
		{
			RoleId:      2,
			Email:       "user2@gmail.com",
			GoogleId:    "asdfghjkl",
			Username:    "user2",
			PhoneNumber: "085123456789",
			Password:    string(hashPasswordUser2),
		},
	}

	return user
}
