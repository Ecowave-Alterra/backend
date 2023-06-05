package profile

import (
	p "github.com/berrylradianh/ecowave-go/helper/password"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (pc *profileUsecase) GetAllUserProfile(user *[]ut.User) error {
	return pc.profileRepo.GetAllUserProfile(user)
}

func (pc *profileUsecase) GetUserProfile(user *ut.User, id int) error {
	return pc.profileRepo.GetUserProfile(user, id)
}

func (pc *profileUsecase) GetUserDetailProfile(userDetail *ut.UserDetail, id int) (bool, error) {
	available, err := pc.profileRepo.GetUserDetailProfile(userDetail, id)
	return available, err
}

func (pc *profileUsecase) UpdateUserProfile(user *ut.User, id int) error {
	return pc.profileRepo.UpdateUserProfile(user, id)
}

func (pc *profileUsecase) UpdateUserDetailProfile(userDetail *ut.UserDetail, id int) error {
	return pc.profileRepo.UpdateUserDetailProfile(userDetail, id)
}

func (pc *profileUsecase) CreateUserDetailProfile(userDetail *ut.UserDetail) error {
	return pc.profileRepo.CreateUserDetailProfile(userDetail)
}

func (pc *profileUsecase) CreateAddressProfile(address *ut.UserAddress) error {
	return pc.profileRepo.CreateAddressProfile(address)
}

func (pc *profileUsecase) GetAllAddressProfile(address *[]ut.UserAddress, idUser int) error {
	return pc.profileRepo.GetAllAddressProfile(address, idUser)
}

func (pc *profileUsecase) GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	return pc.profileRepo.GetAddressByIdProfile(address, idUser, idAddress)
}

func (pc *profileUsecase) UpdateAddressPrimaryProfile(address *ut.UserAddress, idUser int) error {
	return pc.profileRepo.UpdateAddressPrimaryProfile(address, idUser)
}

func (pc *profileUsecase) UpdateAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	return pc.profileRepo.UpdateAddressByIdProfile(address, idUser, idAddress)
}

func (pc *profileUsecase) UpdatePasswordProfile(user *ut.User, oldPassword string, newPassword string, id int) (string, error) {
	if err := p.VerifyPassword(user.Password, oldPassword); err != nil {
		return "Password salah", err
	}

	hashNewPassword, err := p.HashPassword(newPassword)
	if err != nil {
		return "", err
	}

	return "", pc.profileRepo.UpdatePasswordProfile(string(hashNewPassword), id)
}
