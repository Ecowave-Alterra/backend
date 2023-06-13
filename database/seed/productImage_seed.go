package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductImage() []*productEntity.ProductImage {
	product_images := []*productEntity.ProductImage{
		{
			ProductId:       1,
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       1,
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       2,
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       2,
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
	}

	return product_images
}
