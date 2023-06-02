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
	CreateAddressProfile(address *ut.UserAddress) error
	GetAllAddressProfile(address *[]ut.UserAddress, idUser int) error
	GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error
	UpdateAddressProfile(address *ut.UserAddress, idUser int, idAddress int) error
	GetPasswordProfile(user *ut.User, id int) (string, error)
	UpdatePasswordProfile(user *ut.User, oldPassword string, newPassword string, id int) (string, error)
}

type profileUsecase struct {
	profileRepo pr.ProfileRepo
}

func New(profileRepo pr.ProfileRepo) *profileUsecase {
	return &profileUsecase{
		profileRepo,
	}
}
