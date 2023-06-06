package user

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model     `json:"-"`
	FullName        string `json:"FullName" form:"FullName"`
	ProfilePhotoUrl string `json:"ProfilePhotoUrl" form:"ProfilePhotoUrl"`
	EcoPoint        int    `json:"EcoPoint" form:"EcoPoint"`
	UserId          uint   `json:"UserId" form:"UserId"`
}

type User2Response struct {
	UserId          int
	FullName        string
	Email           string
	Username        string
	PhoneNumber     string
	ProfilePhotoUrl string
	UserDetailId    int
}
