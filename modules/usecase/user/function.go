package user

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (uc *userUsecase) CreateUser(user *ut.User) error {

	err := uc.userRepo.GetUserEmail(user.Email)
	if err != nil {
		return err
	}
	res := uc.userRepo.CreateUser(user)
	if res != nil {
		return res
	}
	return nil
}
