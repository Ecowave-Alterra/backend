package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model   `json:"-"`
	RoleId        uint
	Email         string     `json:"Email" form:"Email" validate:"required,email"`
	GoogleId      string     `json:"GoogleId" form:"GoogleId"`
	Username      string     `json:"Username" form:"Username" validate:"required"`
	PhoneNumber   string     `json:"PhoneNumber" form:"PhoneNumber" validate:"required,min=10,max=15"`
	Password      string     `json:"Password" form:"Password" validate:"required"`
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

type UserResponse struct {
	FullName    string
	Username    string
	Email       string
	PhoneNumber string
	EcoPoint    int
}
