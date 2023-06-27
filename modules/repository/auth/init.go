package auth

import (
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type AuthRepo interface {
	GetUserByEmail(email string) (*ue.User, error)
	Login(email string) (*ue.AuthResponse, string, uint, error)
	LoginGoogleId(googleId string) (*ue.AuthResponse, uint, error)
	CreateUser(user *ue.RegisterRequest) error
	CreateUserGoogle(user *ue.RegisterGoogleRequest) error
	UserRecovery(userId uint, codeVer string) error
	UpdateUserRecovery(userId uint, codeVer string) error
	GetUserRecovery(userId uint) (ue.UserRecovery, error)
	ChangePassword(user ue.RecoveryRequest) error
	DeleteUserRecovery(userId uint) error
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
