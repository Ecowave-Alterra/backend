package auth

import (
	pw "github.com/berrylradianh/ecowave-go/helper/password"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

// func (ac *authUsecase) LoginAdmin(admin *ut.User) (string, error) {
// 	passwordDB, err := ac.authRepo.LoginAdmin(admin)
// 	if err != nil {
// 		return "", err
// 	}

// 	if err := password.VerifyPassword(passwordDB, admin.Password); err != nil {
// 		return "", err
// 	}

// 	token, err := jwt.CreateToken(int(admin.ID), admin.Email)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }

func (ac *authUsecase) LoginAdmin(email, password string) (*ut.User, error) {
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
