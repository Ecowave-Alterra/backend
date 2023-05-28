package seed

import (
	productEntity "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func CreateProductImage() []*productEntity.Product_Image {
	product_images := []*productEntity.Product_Image{
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
