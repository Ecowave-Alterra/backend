package ecommerce

import (
	"fmt"
	"math"

	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
)

func (ec *ecommerceUseCase) GetAllProduct(products *[]ep.Product, offset, pageSize int) (*[]ee.ProductResponse, int64, error) {
	var productResponses []ee.ProductResponse
	var productImage ep.ProductImage

	products, count, err := ec.ecommerceRepo.GetAllProduct(products, offset, pageSize)
	if err != nil {
		return &productResponses, count, err
	}

	for _, product := range *products {
		productImages, err := ec.ecommerceRepo.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return &productResponses, count, err
		}

		var productImageURLs []string
		for _, image := range productImages {
			productImageURLs = append(productImageURLs, image.ProductImageUrl)
		}

		avgRating, err := ec.ecommerceRepo.AvgRating(product.ProductID)
		if err != nil {
			return &productResponses, count, err
		}

		reviews, err := ec.ecommerceRepo.GetProductByID(product.ProductID)
		if err != nil {
			return &productResponses, count, err
		}

		productResponse := ee.ProductResponse{
			ProductId:       product.ProductID,
			Name:            product.Name,
			Category:        product.ProductCategory.Category,
			Stock:           int(product.Stock),
			Price:           product.Price,
			Status:          product.Status,
			Description:     product.Description,
			ProductImageUrl: productImageURLs,
			AvgRating:       math.Round(avgRating*10) / 10,
			Review:          reviews,
		}

		productResponses = append(productResponses, productResponse)
	}

	return &productResponses, count, nil
}

// func (ec *ecommerceUseCase) FilterProductByCategoryAndPrice(qCategory, qPrice string, offset, pageSize int) (*[]ee.ProductResponse, int64, error) {
// 	var products *[]ep.Product
// 	var productImage ep.ProductImage
// 	var productResponses []ee.ProductResponse
// 	var total int64
// 	var err error

// 	if qCategory != "" {
// 		if qPrice == "max" {
// 			products, total, err = ec.ecommerceRepo.FilterProductByCategoryAndPriceMax(qCategory, products, offset, pageSize)
// 			if err != nil {
// 				return &productResponses, total, err
// 			}
// 		} else if qPrice == "min" {
// 			products, total, err = ec.ecommerceRepo.FilterProductByCategoryAndPriceMin(qCategory, products, offset, pageSize)
// 			if err != nil {
// 				return &productResponses, total, err
// 			}
// 		} else {
// 			products, total, err = ec.ecommerceRepo.FilterProductByCategory(qCategory, products, offset, pageSize)
// 			if err != nil {
// 				return &productResponses, total, err
// 			}
// 		}
// 	} else {
// 		if qPrice == "max" {
// 			products, total, err = ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMax(products, offset, pageSize)
// 			if err != nil {
// 				return &productResponses, total, err
// 			}
// 		} else if qPrice == "min" {
// 			products, total, err = ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMin(products, offset, pageSize)
// 			if err != nil {
// 				return &productResponses, total, err
// 			}
// 		}
// 	}

// 	for _, product := range *products {
// 		productImages, err := ec.ecommerceRepo.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
// 		if err != nil {
// 			return &productResponses, total, err
// 		}

// 		var imageURL string
// 		for _, image := range productImages {
// 			imageURL = image.ProductImageUrl
// 			break
// 		}

// 		productResponse := ee.ProductResponse{
// 			Name:            product.Name,
// 			Price:           product.Price,
// 			Rating:          product.Rating,
// 			ProductImageUrl: imageURL,
// 		}

// 		productResponses = append(productResponses, productResponse)
// 	}

// 	return &productResponses, total, nil
// }
