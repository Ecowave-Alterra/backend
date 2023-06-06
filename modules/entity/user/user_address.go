package user

import "gorm.io/gorm"

type UserAddress struct {
	*gorm.Model `json:"-"`
	Recipient   string `json:"Recipient" form:"Recipient" validate:"required"`
	PhoneNumber string `json:"PhoneNumber" form:"PhoneNumber" validate:"required,min=10,max=13"`
	Address     string `json:"Address" form:"Address" validate:"required"`
	Note        string `json:"Note" form:"Note"`
	Mark        string `json:"Mark" form:"Mark"`
	IsPrimary   bool   `json:"IsPrimary" form:"IsPrimary"`
	UserId      uint   `json:"UserId" form:"UserId"`
}

type UserAddressResponse struct {
	UserAddress int
	Recipient   string
	PhoneNumber string
	Address     string
	Note        string
	Mark        string
	IsPrimary   bool
	UserId      int
}
