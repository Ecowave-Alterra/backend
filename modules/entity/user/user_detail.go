package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model     `json:"-"`
	Name            string `json:"Name" form:"Name"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
	Point           int    `json:"Point" form:"Point"`
	UserId          uint   `json:"UserId" form:"UserId"`
}
