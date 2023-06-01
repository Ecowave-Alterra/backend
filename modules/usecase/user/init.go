package user

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ur "github.com/berrylradianh/ecowave-go/modules/repository/user"
)

type UserUsecase interface {
	CreateUser(user *ut.UserRequest) error
	LoginUser(user *ut.User) (error, interface{})
}

type userUsecase struct {
	userRepo ur.UserRepo
}

func New(userRepo ur.UserRepo) *userUsecase {
	return &userUsecase{
		userRepo,
	}
}
