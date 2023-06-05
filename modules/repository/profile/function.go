package profile

import (
	"log"

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

func (pr *profileRepo) GetUserDetailProfile(userDetail *ut.UserDetail, id int) (bool, error) {
	result := pr.db.Raw("SELECT * FROM user_details WHERE user_id = ?", id).Scan(&userDetail)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (pr *profileRepo) CreateUserDetailProfile(userDetail *ut.UserDetail) error {
	if err := pr.db.Save(&userDetail).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdateUserProfile(user *ut.User, id int) error {
	if err := pr.db.Raw("UPDATE users SET email = ?, username = ?, phone_number = ? WHERE id = ?", user.Email, user.Username, user.PhoneNumber, id).Scan(&user).Error; err != nil {
		return err
	}

	// if err := pr.db.Where("id = ?", id).Updates(&user).Error; err != nil {
	// 	return err
	// }

	return nil
}

func (pr *profileRepo) UpdateUserDetailProfile(userDetail *ut.UserDetail, id int) error {
	if err := pr.db.Raw("UPDATE user_details SET full_name = ?, profile_photo_url = ? WHERE user_id = ?", userDetail.FullName, userDetail.ProfilePhotoUrl, id).Scan(&userDetail).Error; err != nil {
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

func (pr *profileRepo) GetAllAddressProfile(address *[]ut.UserAddress, idUser int) error {
	if err := pr.db.Where("user_id = ?", idUser).Find(&address).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) GetAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	if err := pr.db.Where("user_id = ? AND id = ?", idUser, idAddress).First(&address).Error; err != nil {
		return err
	}

	return nil
}

func (pr *profileRepo) UpdateAddressPrimaryProfile(address *ut.UserAddress, idUser int) error {
	if err := pr.db.Model(&address).Where("user_id = ?", idUser).Update("is_primary", false).Error; err != nil {
		return err
	}

	if err := pr.db.Raw("UPDATE user_addresses SET is_primary = false WHERE user_id = ?", idUser).Scan(&address).Error; err != nil {
		return err
	}

	log.Println(address.IsPrimary)

	return nil
}

func (pr *profileRepo) UpdateAddressByIdProfile(address *ut.UserAddress, idUser int, idAddress int) error {
	if err := pr.db.Where("user_id = ? AND id = ?", idUser, idAddress).Updates(&address).Error; err != nil {
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
