package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUser() []*ut.User {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")
	hashPasswordAdmin, _ := password.HashPassword("admin123")

	user := []*ut.User{
		{
			Email:    "admin@gmail.com",
			Username: "admin",
			Password: string(hashPasswordAdmin),
			RoleId:   1,
		},
		{
			Email:    "user1@gmail.com",
			Username: "user1",
			Password: string(hashPasswordUser1),
			RoleId:   2,
		},
		{
			Email:    "user2@gmail.com",
			Username: "user2",
			Password: string(hashPasswordUser2),
			RoleId:   2,
		},
	}

	return user
}
