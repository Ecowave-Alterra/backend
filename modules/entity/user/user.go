package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*gorm.Model `json:"-"`
	RoleId      uint
	Email       string `json:"Email" form:"Email" validate:"required,email"`
	GoogleId    string `json:"GoogleId" form:"GoogleId"`
	Username    string `json:"Username" form:"Username" validate:"required"`
	PhoneNumber string `json:"PhoneNumber" form:"PhoneNumber" validate:"required"`
	Password    string `json:"Password" form:"Password" validate:"required"`
}

type UserResponseLogin struct {
	Email    string
	Username string
	Token    string
}
