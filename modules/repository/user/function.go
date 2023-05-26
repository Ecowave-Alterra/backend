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
func (r *userRepo) CreateUser(user *ut.User) error {

	err := r.db.Save(&user).Error
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
