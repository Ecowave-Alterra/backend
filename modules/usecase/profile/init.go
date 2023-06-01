package profile

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	pr "github.com/berrylradianh/ecowave-go/modules/repository/profile"
)

type ProfileUsecase interface {
	GetUserProfile(id int) (*ut.UserResponse, error)
	GetUser2Profile(id int) (*ut.User2Response, *ut.User, *ut.UserDetail, error)
	UpdateUserProfile(user *ut.User, id int) error
	UpdateUserDetailProfile(userDetail *ut.UserDetail, id int) error
	// GetUserDetailProfile(userDetail *ut.UserDetail, id int) error
}

type profileUsecase struct {
	profileRepo pr.ProfileRepo
}

func New(profileRepo pr.ProfileRepo) *profileUsecase {
	return &profileUsecase{
		profileRepo,
	}
}
