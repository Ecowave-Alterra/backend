package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProduct() []*productEntity.Product {
	products := []*productEntity.Product{
		{
			Product_category_id:    1,
			Product_description_id: 1,
			Name:                   "Product Name 1",
			Stock:                  10,
			Price:                  30000,
			Status:                 "tersedia",
			Rating:                 0.00,
		},
		{
			Product_category_id:    1,
			Product_description_id: 2,
			Name:                   "Product Name 2",
			Stock:                  0,
			Price:                  36000,
			Status:                 "habis",
			Rating:                 0.00,
		},
	}

	return products
}
