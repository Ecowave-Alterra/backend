package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductImage() []*productEntity.ProductImage {
	product_images := []*productEntity.ProductImage{
		{
			ProductId:       "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "f71ff306-ebd7-45e5-9607-5b908dd1c423",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "ba08266c-0926-484e-a610-c2a48ac6179d",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "ba08266c-0926-484e-a610-c2a48ac6179d",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "a3325f33-e01a-4e40-9ca7-5d84c4337094",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
	}

	return product_images
}
