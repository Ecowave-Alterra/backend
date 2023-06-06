package auth

import (
	"errors"

	pw "github.com/berrylradianh/ecowave-go/helper/password"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	"github.com/berrylradianh/ecowave-go/middleware/jwt"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ac *authUsecase) Register(request *ue.RegisterRequest) error {
	if err := vld.ValidateRegister(request); err != nil {
		return err
	}

	_, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exist")
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = string(hashedPassword)

	err = ac.authRepo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(request *ue.LoginRequest) (*ue.User, string, error) {
	if err := vld.ValidateLogin(request); err != nil {
		return nil, "", err
	}

	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, "", errors.New("Email atau password salah")
	}

	err = pw.VerifyPassword(user.Password, request.Password)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, "", errors.New("Email atau password salah")
	}

	token, err := jwt.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
