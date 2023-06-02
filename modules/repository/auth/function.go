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

func (ar *authRepo) CreateUser(user *ue.User) error {
	existingUser := &ue.User{}
	err := ar.db.Where("email = ?", user.Email).First(existingUser).Error
	if err != nil {
		err = ar.db.Create(&user).Error
		if err != nil {
			return err
		}
	} else {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	return nil
}
