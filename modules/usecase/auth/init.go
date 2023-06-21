package auth

import (
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ar "github.com/berrylradianh/ecowave-go/modules/repository/auth"
)

type AuthUsecase interface {
	Login(request *ue.LoginRequest) (*ue.User, string, error)
	Register(user *ue.RegisterRequest) error
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
