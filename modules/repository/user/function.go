package user

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (r *userRepo) GetUserEmail(email string) error {
	var user ut.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return err
	}
	return nil
}
func (r *userRepo) CreateUser(user *ut.UserRequest) error {

	userTable := ut.User{
		RoleId:   2,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}

	userDetail := ut.UserDetail{
		Name:  user.Name,
		Phone: user.PhoneNumber,
	}

	//save table user
	err := r.db.Save(&userTable).Error
	if err != nil {
		return err
	}
	// save table user detail
	err = r.db.Save(&userDetail).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *userRepo) LoginUser(user *ut.User) (error, string) {

	err := r.db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return err, ""
	}

	return nil, user.Password
}
