package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateUser() []*ue.User {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")
	hashPasswordAdmin, _ := password.HashPassword("admin123")
	users := []*ue.User{
		{
			Email:    "admin@gmail.com",
			Password: string(hashPasswordAdmin),
			RoleId:   1,
		},
		{
			Email:    "user1@gmail.com",
			Password: string(hashPasswordUser1),
			RoleId:   2,
		},
		{
			Email:    "user2@gmail.com",
			Password: string(hashPasswordUser2),
			RoleId:   2,
		},
	}
	return users
}
