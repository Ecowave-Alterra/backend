package auth

import ac "github.com/berrylradianh/ecowave-go/modules/usecase/auth"

type AuthHandler struct {
	authUsecase ac.AuthUsecase
}

func New(authUsecase ac.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase,
	}
}
