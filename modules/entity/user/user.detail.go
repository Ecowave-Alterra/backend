package user

import (
	"github.com/jinzhu/gorm"
)

type UserDetail struct {
	*gorm.Model     `json:"-"`
	UserId          uint   `json:"userId" form:"userId" validate:"required"`
	Name            string `json:"Name" form:"Name" validate:"required"`
	Point           uint   `json:"Point" form:"Point"`
	Phone           string `json:"Phone" form:"Phone" validate:"required,min=10,max=15"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
}
