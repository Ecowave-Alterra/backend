package seed

import (
	"time"

	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProduct() []*productEntity.Product {
	products := []*productEntity.Product{
		{
			ProductId:         "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductCategoryId: 1,
			Name:              "Product Name 1",
			Stock:             10,
			Price:             30000,
			Status:            "tersedia",
			Rating:            0.00,
			Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         time.Now(),
		},
		{
			ProductId:         "06dcbb14-507e-41db-8942-a0e93f50ebe5",
			ProductCategoryId: 2,
			Name:              "Product Name 2",
			Stock:             0,
			Price:             36000,
			Status:            "habis",
			Rating:            0.00,
			Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			CreatedAt:         time.Now(),
		},
	}

	return products
}
