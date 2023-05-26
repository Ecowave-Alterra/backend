package user

import (
	p "github.com/berrylradianh/ecowave-go/helper/password"
	jwt "github.com/berrylradianh/ecowave-go/middleware/jwt"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (uc *userUsecase) CreateUser(user *ut.User) error {

	err := uc.userRepo.GetUserEmail(user.Email)
	if err != nil {
		return err
	}
	user.RoleId = 2

	password := user.Password
	hash, err := p.HashPassword(password)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	res := uc.userRepo.CreateUser(user)
	if res != nil {
		return res
	}
	return nil
}
func (uc *userUsecase) LoginUser(user *ut.User) (error, interface{}) {

	err, hashPassword := uc.userRepo.LoginUser(user)
	if err != nil {
		return err, nil
	}

	password := user.Password
	err = p.VerifyPassword(hashPassword, password)
	if err != nil {
		return err, nil
	}

	token, err := jwt.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return err, nil
	}

	loginResponse := ut.UserResponseLogin{Username: user.Username, Email: user.Email, Token: token}

	return nil, loginResponse

}
