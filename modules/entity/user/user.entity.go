package user

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model   `json:"-"`
	RoleId        uint
	Email         string        `json:"Email" form:"Email" validate:"required,email"`
	GoogleId      string        `json:"GoogleId" form:"GoogleId"`
	Username      string        `json:"Username" form:"Username" validate:"required"`
	Password      string        `json:"Password" form:"Password" validate:"required,min=8"`
	UserDetail    UserDetail    `gorm:"foreignKey:UserId"`
	UserAddresses []UserAddress `gorm:"foreignKey:UserId"`
	UserRecovery  UserRecovery  `gorm:"foreignKey:UserId" json:"-"`
}

type RegisterRequest struct {
	Name     string `json:"Name" form:"Name" validate:"required"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Username string `json:"Username" form:"Username" validate:"required"`
	Phone    string `json:"Phone" form:"Phone" validate:"required,min=10,max=15,numeric"`
	Password string `json:"Password" form:"Password" validate:"required,min=8"`
}
type RegisterGoogleRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	GoogleId string `json:"GoogleId" form:"GoogleId"`
	Name     string `json:"Name" form:"Name" validate:"required"`
	Username string `json:"Username" form:"Username" validate:"required"`
	Phone    string `json:"Phone" form:"Phone" validate:"required,min=10,max=15,numeric"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
type LoginGoogleRequest struct {
	GoogleId string `json:"GoogleId" form:"GoogleId"`
}

type AuthResponse struct {
	ID            uint   `json:"Id" form:"Id"`
	GoogleId      string `json:"GoogleId" form:"GoogleId"`
	Email         string `json:"Email" form:"Email" validate:"required,email"`
	Username      string `json:"Username" form:"Username" validate:"required"`
	Name          string `json:"Name" form:"Name"`
	Point         uint   `json:"Point" form:"Point"`
	Phone         string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	ProfilePhoto  string `json:"ProfilePhoto" form:"ProfilePhoto"`
	UserAddresses UserAddress
	Token         string `json:"Token" form:"Token"`
}

type UserLogin struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type UserResponseLogin struct {
	Email    string
	Username string
	Token    string
}

type UserRequest struct {
	Email    string `json:"Email" form:"Email" validate:"email"`
	Username string `json:"Username" form:"Username"`
}

type UserResponse struct {
	Id           uint   `json:"Id"`
	GoogleId     string `json:"GoogleId"`
	RoleId       uint   `json:"RoleId"`
	Name         string `json:"Name"`
	Username     string `json:"Username"`
	Email        string `json:"Email"`
	Phone        string `json:"Phone"`
	Point        uint   `json:"Point"`
	ProfilePhoto string `json:"ProfilePhoto"`
	Addresses    []UserAddressResponse
}

type UserPasswordRequest struct {
	OldPassword        string `json:"OldPassword" form:"OldPassword" validate:"required"`
	Password           string `json:"Password" form:"Password" validate:"required"`
	ConfirmNewPassword string `json:"ConfirmNewPassword" form:"ConfirmNewPassword" validate:"required"`
}
