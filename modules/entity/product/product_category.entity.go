package product

import "gorm.io/gorm"

type ProductCategory struct {
	*gorm.Model `json:"-"`

	Name string `json:"Name" form:"Name" validate:"required"`
}
