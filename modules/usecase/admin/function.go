package admin

import (
	"github.com/berrylradianh/ecowave-go/helper/password"
	jwt "github.com/berrylradianh/ecowave-go/middleware/jwt"
	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
)

func (ac *adminUsecase) LoginAdmin(admin *at.Admin) (string, error) {
	passwordDB, err := ac.adminRepo.LoginAdmin(admin)
	if err != nil {
		return "", err
	}

	if err := password.VerifyPassword(passwordDB, admin.Password); err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(int(admin.ID), admin.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
