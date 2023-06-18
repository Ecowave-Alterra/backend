package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model     `json:"-"`
	Name            string `json:"Name" form:"Name"`
	Point           int    `json:"Point" form:"Point"`
	Phone           string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
	UserId          uint   `json:"UserId" form:"UserId"`
}

type UserDetailRequest struct {
	Name            string `json:"Name" form:"Name"`
	Phone           string `json:"Phone" form:"Phone" validate:"min=10,max=13"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
}
