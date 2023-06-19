package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductImage() []*productEntity.ProductImage {
	product_images := []*productEntity.ProductImage{
		{
			ProductId:       "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "06dcbb14-507e-41db-8942-a0e93f50ebe5",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "06dcbb14-507e-41db-8942-a0e93f50ebe5",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
	}

	return product_images
}
