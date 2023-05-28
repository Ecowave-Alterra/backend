package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductCategory() []*productEntity.Product_Category {
	productCategories := []*productEntity.Product_Category{
		{
			Category: "perabot",
		},
		{
			Category: "kantong",
		},
	}

	return productCategories
}
