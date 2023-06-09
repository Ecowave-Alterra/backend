package product

import "gorm.io/gorm"

type ProductCategory struct {
	*gorm.Model `json:"-"`

	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"Name" form:"Name" validate:"required"`
}
