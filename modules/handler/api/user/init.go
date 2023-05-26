package user

import (
	uc "github.com/berrylradianh/ecowave-go/modules/usecase/user"
)

type UserHandler struct {
	userUC uc.UserUsecase
}

func New(userUC uc.UserUsecase) *UserHandler {
	return &UserHandler{
		userUC,
	}
}
