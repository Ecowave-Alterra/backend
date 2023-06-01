package user

import "gorm.io/gorm"

type UserAddress struct {
	*gorm.Model `json:"-"`
	Recipient   string `json:"Recipient" form:"Recipient"`
	PhoneNumber string `json:"PhoneNumber" form:"PhoneNumber"`
	Address     string `json:"Address" form:"Address"`
	Note        string `json:"Note" form:"Note"`
	Mark        string `json:"Mark" form:"Mark"`
	IsPrimary   bool   `json:"IsPrimary" form:"IsPrimary"`
	UserId      uint   `json:"UserId" form:"UserId"`
}
