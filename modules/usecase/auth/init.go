package auth

import (
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ar "github.com/berrylradianh/ecowave-go/modules/repository/auth"
)

type AuthUsecase interface {
	LoginAdmin(email, password string) (*ue.User, error)
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(adminRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		adminRepo,
	}
}
