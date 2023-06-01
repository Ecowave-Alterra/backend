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
	Password    string `json:"Password" form:"Password" validate:"required"`
}
type UserRequest struct {
	Name        string `json:"Name" form:"Name" validate:"required"`
	Email       string `json:"Email" form:"Email" validate:"required,email"`
	Username    string `json:"Username" form:"Username" validate:"required"`
	PhoneNumber string `json:"PhoneNumber" form:"PhoneNumber" validate:"required,min=10,max=15"`
	Password    string `json:"Password" form:"Password" validate:"required"`
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
