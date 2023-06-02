package auth

import (
	pw "github.com/berrylradianh/ecowave-go/helper/password"
	"github.com/berrylradianh/ecowave-go/middleware/jwt"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ac *authUsecase) Register(user *ue.User) error {
	hashedPassword, err := pw.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = ac.authRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(email, password string) (*ue.User, string, error) {
	user, err := ac.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	err = pw.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, "", err
	}

	token, err := jwt.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
