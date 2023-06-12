package seed

import (
	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductCategory() *[]pct.ProductCategory {
	return &[]pct.ProductCategory{
		{
			Category: "perabot",
		},
		{
			Category: "kantong",
		},
	}
}
