package profile

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	pr "github.com/berrylradianh/ecowave-go/modules/repository/user/profile"
)

type ProfileUsecase interface {
	GetAllUserProfile(user *[]ut.User) error
	GetUserProfile(user *ut.User, id int) error
	GetUserDetailProfile(userDetail *ut.UserDetail, id int) error
	UpdateUserProfile(userRequest *ut.UserRequest, id int) error
	UpdateUserDetailProfile(userDetailRequest *ut.UserDetailRequest, id int) error

	CreateAddressProfile(address *ut.UserAddress) error
	GetAllAddressProfileNoPagination(address *[]ut.UserAddress, idUser int) error
	GetAllAddressProfile(address *[]ut.UserAddress, idUser, offset, pageSize int) (*[]ut.UserAddress, int64, error)
	GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error
	UpdateAddressPrimaryProfile(address *ut.UserAddress, idUser int) error
	UpdateAddressByIdProfile(addressRequest *ut.UserAddressRequest, idUser int, idAddress int) error

	UpdatePasswordProfile(user *ut.User, oldPassword string, newPassword string, id int) (string, error)

	GetAllProvince() ([]ut.Province, error)
	GetAllCityByProvince(provinceId string) ([]ut.City, error)
}

type profileUsecase struct {
	profileRepo pr.ProfileRepo
}

func New(profileRepo pr.ProfileRepo) *profileUsecase {
	return &profileUsecase{
		profileRepo,
	}
}
