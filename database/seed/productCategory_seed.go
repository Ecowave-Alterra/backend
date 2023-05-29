package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductCategory() []*productEntity.ProductCategory {
	productCategories := []*productEntity.ProductCategory{
		{
			Category: "perabot",
		},
		{
			Category: "kantong",
		},
	}

	return productCategories
}
