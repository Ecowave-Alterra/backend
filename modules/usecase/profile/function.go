package profile

import ut "github.com/berrylradianh/ecowave-go/modules/entity/user"

func (pc *profileUsecase) GetUserProfile(id int) (*ut.UserResponse, error) {
	var userDB *ut.User
	var userDetailDB *ut.UserDetail

	user, err := pc.profileRepo.GetUserProfile(userDB, id)
	if err != nil {
		return nil, err
	}

	userDetail, err := pc.profileRepo.GetUserDetailProfile(userDetailDB, id)
	if err != nil {
		return nil, err
	}

	userProfileResponse := &ut.UserResponse{
		FullName:    userDetail.FullName,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		EcoPoint:    userDetail.EcoPoint,
	}

	return userProfileResponse, nil
}

func (pc *profileUsecase) GetUser2Profile(id int) (*ut.User2Response, *ut.User, *ut.UserDetail, error) {
	var userDB *ut.User
	var userDetailDB *ut.UserDetail

	user, err := pc.profileRepo.GetUserProfile(userDB, id)
	if err != nil {
		return nil, nil, nil, err
	}

	userDetail, err := pc.profileRepo.GetUserDetailProfile(userDetailDB, id)
	if err != nil {
		return nil, nil, nil, err
	}

	userProfileResponse := &ut.User2Response{
		FullName:        userDetail.FullName,
		Username:        user.Username,
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		ProfilePhotoUrl: userDetail.ProfilePhotoUrl,
	}

	return userProfileResponse, user, userDetail, nil
}

func (pc *profileUsecase) UpdateUserProfile(user *ut.User, id int) error {
	return pc.profileRepo.UpdateUserProfile(user, id)
}

func (pc *profileUsecase) UpdateUserDetailProfile(userDetail *ut.UserDetail, id int) error {
	return pc.profileRepo.UpdateUserDetailProfile(userDetail, id)
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

func (pc *profileUsecase) UpdateAddressProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	return pc.profileRepo.UpdateAddressProfile(address, idUser, idAddress)
}
