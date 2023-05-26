package user

import (
	p "github.com/berrylradianh/ecowave-go/helper/password"
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

	user.RoleId = 2

	password := user.Password
	hash, err := p.HashPassword(password)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	save := r.db.Save(&user)
	if save != nil {
		return save.Error
	}

	return nil
}
