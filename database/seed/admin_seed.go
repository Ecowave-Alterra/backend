package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateAdmin() *ut.User {
	passwordAdmin := "admin123"
	hashPasswordAdmin, _ := password.HashPassword(passwordAdmin)
	admin := &ut.User{
		Email:    "admin@gmail.com",
		Password: string(hashPasswordAdmin),
		RoleId:   1,
	}

	return admin
}
