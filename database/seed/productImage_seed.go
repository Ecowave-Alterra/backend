package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductImage() []*productEntity.ProductImage {
	product_images := []*productEntity.ProductImage{
		{
			Product_id:        1,
			Product_image_url: "https://picsum.photos/200/300",
		},
		{
			Product_id:        1,
			Product_image_url: "https://picsum.photos/200/300",
		},
		{
			Product_id:        2,
			Product_image_url: "https://picsum.photos/200/300",
		},
		{
			Product_id:        2,
			Product_image_url: "https://picsum.photos/200/300",
		},
	}

	return product_images
}
