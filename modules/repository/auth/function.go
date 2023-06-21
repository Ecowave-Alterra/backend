package auth

import (
	"errors"

	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ar *authRepo) GetUserByEmail(email string) (*ue.User, error) {
	user := &ue.User{}
	err := ar.db.Where("email = ?", email).First(&user).Error
	if err != nil {
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
func (ar *authRepo) GetUserRecovery(userId uint) (ue.UserRecovery, error) {
	var recovery ue.UserRecovery
	err := ar.db.Where("user_id = ?", userId).First(&recovery).Error
	if err != nil {
		return recovery, errors.New("Record Not Found")
	}

	return recovery, nil
}

func (ar *authRepo) UserRecovery(userId uint, codeVer string) error {

	userRecover := ue.UserRecovery{
		UserId: userId,
		Code:   codeVer,
	}
	if err := ar.db.Create(&userRecover).Error; err != nil {
		return err
	}

	return nil
}
func (ar *authRepo) UpdateUserRecovery(userId uint, codeVer string) error {

	userRecover := ue.UserRecovery{
		UserId: userId,
		Code:   codeVer,
	}
	result := ar.db.Model(&ue.UserRecovery{}).Where("user_id = ?", userId).Updates(&userRecover)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
func (ar *authRepo) ChangePassword(user ue.RecoveryRequest) error {

	result := ar.db.Model(&ue.User{}).Where("email = ?", user.Email).Update("password", user.Password)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
func (ar *authRepo) DeleteUserRecovery(userId uint) error {

	var userRecovery ue.UserRecovery
	result := ar.db.Where("user_id = ?", userId).Delete(&userRecovery)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("nothing has changed")
	}

	return nil
}
