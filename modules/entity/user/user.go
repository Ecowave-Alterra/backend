package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model   `json:"-"`
	RoleId        uint
	Email         string     `json:"Email" form:"Email" validate:"required,email"`
	GoogleId      string     `json:"GoogleId" form:"GoogleId"`
	Username      string     `json:"Username" form:"Username" validate:"required"`
	Password      string     `json:"Password" form:"Password" validate:"required,min=8"`
	UserDetail    UserDetail `gorm:"foreignKey:UserId"`
	UserAddresses []UserAddress
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
	Id              int    `json:"Id"`
	GoogleId        string `json:"GoogleId"`
	RoleId          int    `json:"RoleId"`
	Name            string `json:"Name"`
	Username        string `json:"Username"`
	Email           string `json:"Email"`
	Phone           string `json:"Phone"`
	Point           int    `json:"Point"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl"`
	Addresses       []UserAddressResponse
}

type UserPasswordRequest struct {
	OldPassword        string `json:"OldPassword" form:"OldPassword" validate:"required"`
	Password           string `json:"Password" form:"Password" validate:"required"`
	ConfirmNewPassword string `json:"ConfirmNewPassword" form:"ConfirmNewPassword" validate:"required"`
}
