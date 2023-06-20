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
			Stock:             10,
			Weight:            3000,
			Price:             30000,
			Status:            "tersedia",
			Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
			DeletedAt:         nil,
		},
		{
			ProductId:         "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductCategoryId: 2,
			Name:              "Product Name 2",
			Stock:             100,
			Weight:            4000,
			Price:             36000,
			Status:            "habis",
			Rating:            0.00,
			Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
			DeletedAt:         nil,
		},
		{
			ProductId:         "ba08266c-0926-484e-a610-c2a48ac6179d",
			ProductCategoryId: 3,
			Name:              "Product Name 3",
			Stock:             100,
			Weight:            5000,
			Price:             30000,
			Status:            "tersedia",
			Rating:            0.00,
			Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			CreatedAt:         time.Now(),
			UpdatedAt:         nil,
			DeletedAt:         nil,
		},
	}

	return products
}
