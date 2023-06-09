package seed

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"

	"gorm.io/gorm"
)

func CreateProductCategory(db *gorm.DB) *[]pct.ProductCategory {
	return &[]pct.ProductCategory{
		{
			Name: "perabot",
		},
		{
			Name: "kantong",
		},
	}
}
