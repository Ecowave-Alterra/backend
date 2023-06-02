package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model  `json:"-"`
	Name         string `json:"Name" form:"Name" validate:"required"`
	Point        uint   `json:"Point" form:"Point"`
	Phone        string `json:"Phone" form:"Phone" validate:"required,numeric"`
	ProfilePhoto string `json:"ProfilePhoto" form:"ProfilePhoto"`
	UserId       uint   `json:"UserId" form:"UserId"`
}
