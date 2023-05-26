package user

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *ut.User) error
	GetUserEmail(email string) error
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}
