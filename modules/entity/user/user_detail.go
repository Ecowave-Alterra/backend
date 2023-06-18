package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model     `json:"-"`
	Name            string `json:"Name" form:"Name"`
	Point           int    `json:"Point" form:"Point"`
	Phone           string `json:"Phone" form:"Phone" validate:"required,min=10,max=15"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
	UserId          uint   `json:"UserId" form:"UserId"`
}
