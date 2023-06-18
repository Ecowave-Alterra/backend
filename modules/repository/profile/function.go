package profile

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (pr *profileRepo) GetAllUserProfile(user *[]ut.User) error {
	if err := pr.db.Find(&user).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) GetUserProfile(user *ut.User, id int) error {
	if err := pr.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) GetUserDetailProfile(userDetail *ut.UserDetail, id int) error {
	if err := pr.db.Raw("SELECT * FROM user_details WHERE user_id = ?", id).Scan(&userDetail).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) CreateUserDetailProfile(userDetail *ut.UserDetail) error {
	if err := pr.db.Save(&userDetail).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdateUserProfile(userRequest *ut.UserRequest, id int) error {
	if err := pr.db.Raw("UPDATE users SET email = ?, username = ? WHERE id = ?", userRequest.Email, userRequest.Username, id).Scan(&userRequest).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdateUserDetailProfile(userDetailRequest *ut.UserDetailRequest, id int) error {
	if err := pr.db.Raw("UPDATE user_details SET name = ?, phone = ?, profile_photo_url = ? WHERE user_id = ?", userDetailRequest.Name, userDetailRequest.Phone, userDetailRequest.ProfilePhotoUrl, id).Scan(&userDetailRequest).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) CreateAddressProfile(address *ut.UserAddress) error {
	if err := pr.db.Save(&address).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) GetAllAddressProfileNoPagination(address *[]ut.UserAddress, idUser int) error {
	if err := pr.db.Where("user_id = ?", idUser).Find(&address).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) GetAllAddressProfile(address *[]ut.UserAddress, idUser, offset, pageSize int) (*[]ut.UserAddress, int64, error) {
	var count int64

	if err := pr.db.Model(&address).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := pr.db.Offset(offset).Limit(pageSize).Where("user_id = ?", idUser).Find(&address).Error; err != nil {
		return nil, 0, err
	}

	return address, count, nil
}

func (pr *profileRepo) GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	if err := pr.db.Where("user_id = ? AND id = ?", idUser, idAddress).First(&address).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdateAddressPrimaryProfile(address *ut.UserAddress, idUser int) error {
	if err := pr.db.Raw("UPDATE user_addresses SET is_primary = false WHERE user_id = ?", idUser).Scan(&address).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdateAddressByIdProfile(addressRequest *ut.UserAddressRequest, idUser int, idAddress int) error {
	var address ut.UserAddress

	if err := pr.db.Model(address).Where("user_id = ? AND id = ?", idUser, idAddress).Updates(&addressRequest).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdatePasswordProfile(newPassword string, id int) error {
	var user *ut.User
	if err := pr.db.Model(&user).Where("id = ?", id).Update("password", newPassword).Error; err != nil {
		return err
	}

	return nil
}
