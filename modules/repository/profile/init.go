package profile

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type ProfileRepo interface {
	GetUserProfile(user *ut.User, id int) (*ut.User, error)
	GetUserDetailProfile(userDetail *ut.UserDetail, id int) (*ut.UserDetail, error)
	UpdateUserProfile(user *ut.User, id int) error
	UpdateUserDetailProfile(userDetail *ut.UserDetail, id int) error
	CreateAddressProfile(address *ut.UserAddress) error
	GetAllAddressProfile(address *[]ut.UserAddress, idUser int) error
	GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error
	UpdateAddressProfile(address *ut.UserAddress, idUser int, idAddress int) error
}

type profileRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ProfileRepo {
	return &profileRepo{
		db,
	}
}
