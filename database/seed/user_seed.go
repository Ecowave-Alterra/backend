package seed

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
	"gorm.io/gorm"
)

func CreateAdmin(db *gorm.DB) *at.Admin {
	emailAdmin := "admin@gmail.com"
	passwordAdmin := "admin123"

	hashPasswordAdmin, _ := password.HashPassword(passwordAdmin)

	return &at.Admin{
		Email:    emailAdmin,
		Password: string(hashPasswordAdmin),
	}
}
