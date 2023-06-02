package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model `json:"-"`
	Email       string `json:"Email" form:"Email" validate:"required,email"`
	Username    string `json:"Username" form:"Username" validate:"required"`
	GoogleId    string `json:"GoogleId" form:"GoogleId"`
	Password    string `json:"Password" form:"Password" validate:"required"`
	RoleId      uint   `json:"RoleId" form:"RoleId"`
	UserDetail  UserDetail
}
