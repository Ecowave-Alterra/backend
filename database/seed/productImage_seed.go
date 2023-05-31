package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductImage() []*productEntity.ProductImage {
	product_images := []*productEntity.ProductImage{
		{
			ProductId:       1,
			ProductImageUrl: "https://picsum.photos/200/300",
		},
		{
			ProductId:       1,
			ProductImageUrl: "https://picsum.photos/200/300",
		},
		{
			ProductId:       2,
			ProductImageUrl: "https://picsum.photos/200/300",
		},
		{
			ProductId:       2,
			ProductImageUrl: "https://picsum.photos/200/300",
		},
	}

	return product_images
}
