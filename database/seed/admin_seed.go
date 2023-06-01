package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func CreateAdmin() *ue.User {
	passwordAdmin := "admin123"
	hashPasswordAdmin, _ := password.HashPassword(passwordAdmin)
	admin := &ue.User{
		Email:    "admin@gmail.com",
		Password: string(hashPasswordAdmin),
		RoleId:   1,
	}

	return admin
}
