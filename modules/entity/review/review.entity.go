package review

import "gorm.io/gorm"

type Review struct {
	*gorm.Model

	Comment      string `json:"comment" form:"comment"`
	CommentAdmin string `json:"commentAdmin" form:"commentAdmin"`
	PhotoUrl     string `json:"photoUrl" form:"photoUrl"`
	VideoUrl     string `json:"videoUrl" form:"photoUrl"`
}
