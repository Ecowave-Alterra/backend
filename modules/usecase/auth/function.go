package auth

import (
	pw "github.com/berrylradianh/ecowave-go/helper/password"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ac *authUsecase) LoginAdmin(email, password string) (*eu.User, error) {
	user, err := ac.authRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = pw.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
