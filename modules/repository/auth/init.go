package auth

import (
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type AuthRepo interface {
	GetUserByEmail(email string) (*ue.User, error)
	CreateUser(user *ue.RegisterRequest) error
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
