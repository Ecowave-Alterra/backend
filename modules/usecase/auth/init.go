package auth

import (
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ar "github.com/berrylradianh/ecowave-go/modules/repository/auth"
)

type AuthUsecase interface {
	Login(request *ue.LoginRequest) (interface{}, uint, error)
	LoginGoogle(request *ue.LoginGoogleRequest) (interface{}, uint, error)
	Register(user *ue.RegisterRequest) error
	RegisterGoogle(user *ue.RegisterGoogleRequest) error
	ForgotPassword(request ue.ForgotPassRequest) (string, error)
	VerifOtp(request ue.VerifOtp) error
	ChangePassword(request ue.RecoveryRequest) error
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(adminRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		adminRepo,
	}
}
