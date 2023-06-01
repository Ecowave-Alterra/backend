package auth

import (
	"errors"

	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ar *authRepo) GetUserByEmail(email string) (*ut.User, error) {
	user := &ut.User{}
	err := ar.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}
