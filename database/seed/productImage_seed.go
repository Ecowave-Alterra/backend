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
			ProductId:       "859679ad-888b-41ac-a663-4a619acc4ae3",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
		{
			ProductId:       "859679ad-888b-41ac-a663-4a619acc4ae3",
			ProductImageUrl: "https://storage.googleapis.com/ecowave/img/products/bottle.png",
		},
	}

	return product_images
}
