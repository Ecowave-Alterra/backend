package auth

import (
	"errors"

	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ar *authRepo) GetUserByEmail(email string) (*ue.User, error) {
	user := &ue.User{}
	err := ar.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}

func (ar *authRepo) CreateUser(user *ue.RegisterRequest) error {
	existingUser := ue.User{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	userTable := ue.User{
		RoleId:   2,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		UserDetail: ue.UserDetail{
			Name:  user.Name,
			Phone: user.Phone,
		},
	}

	if err := ar.db.Create(&userTable).Error; err != nil {
		return err
	}
	return nil
}
