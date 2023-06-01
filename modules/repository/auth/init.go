package auth

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type AuthRepo interface {
	GetUserByEmail(email string) (*ut.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
