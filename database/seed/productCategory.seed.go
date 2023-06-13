package seed

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductCategory() *[]pe.ProductCategory {
	return &[]pe.ProductCategory{
		{
			Category: "perabot",
		},
		{
			Category: "kantong",
		},
	}
}
