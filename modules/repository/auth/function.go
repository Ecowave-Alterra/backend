package auth

import (
	"errors"

	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

// func (ar *authRepo) LoginAdmin(user *ut.User) (string, error) {
// 	var result *ut.User

// 	if err := ar.db.Where("email = ?", user.Email).First(&result).Error; err != nil {
// 		return "", err
// 	}

// 	return result.Password, nil
// }

func (ar *authRepo) GetUserByEmail(email string) (*ut.User, error) {
	user := &ut.User{}
	err := ar.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}
