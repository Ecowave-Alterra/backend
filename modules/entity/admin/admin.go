package admin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
