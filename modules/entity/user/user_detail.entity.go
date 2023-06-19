package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model     `json:"-"`
	Name            string `json:"Name" form:"Name" validate:"required"`
	Point           uint   `json:"Point" form:"Point"`
	Phone           string `json:"Phone" form:"Phone" validate:"required,numeric"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
	UserId          uint   `json:"UserId" form:"UserId"`
}
