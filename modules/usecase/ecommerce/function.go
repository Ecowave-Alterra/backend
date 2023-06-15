package ecommerce

import (
	"fmt"

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

		var imageURL string
		for _, image := range productImages {
			imageURL = image.ProductImageUrl
			break
		}

		productResponse := ee.ProductResponse{
			Name:            product.Name,
			Price:           product.Price,
			Rating:          product.Rating,
			ProductImageUrl: imageURL,
		}

		productResponses = append(productResponses, productResponse)
	}

	return &productResponses, count, nil
}

func (ec *ecommerceUseCase) GetProductByID(productId string) (ee.ProductDetailResponse, error) {
	var productDetailResponse ee.ProductDetailResponse
	var reviews []ee.ReviewResponse
	var productImage ep.ProductImage

	queryResponse, err := ec.ecommerceRepo.GetProductByID(productId)
	if err != nil {
		return productDetailResponse, err
	}

	for _, value := range queryResponse {
		review := ee.ReviewResponse{
			FullName:        value.FullName,
			ProfilePhotoUrl: value.ProfilePhotoUrl,
			Rating:          value.Rating,
			Comment:         value.Comment,
			CommentAdmin:    value.CommentAdmin,
			PhotoURL:        value.PhotoURL,
			VideoURL:        value.VideoURL,
		}
		reviews = append(reviews, review)
	}

	productImages, err := ec.ecommerceRepo.GetProductImageURLById(fmt.Sprint(queryResponse[0].Id), &productImage)
	if err != nil {
		return productDetailResponse, err
	}

	var productImageURLs []string
	for _, image := range productImages {
		productImageURLs = append(productImageURLs, image.ProductImageUrl)
	}

	avgRating, err := ec.ecommerceRepo.AvgRating(productId)
	if err != nil {
		return productDetailResponse, err
	}

	productDetailResponse = ee.ProductDetailResponse{
		ProductId:       queryResponse[0].ProductId,
		Name:            queryResponse[0].Name,
		Category:        queryResponse[0].Category,
		Stock:           queryResponse[0].Stock,
		Price:           queryResponse[0].Price,
		Status:          queryResponse[0].Status,
		Description:     queryResponse[0].Description,
		ProductImageUrl: productImageURLs,
		AvgRating:       avgRating,
		Review:          reviews,
	}

	return productDetailResponse, nil
}

func (ec *ecommerceUseCase) FilterProductByCategoryAndPrice(qCategory, qPrice string, offset, pageSize int) (*[]ee.ProductResponse, int64, error) {
	var products *[]ep.Product
	var productImage ep.ProductImage
	var productResponses []ee.ProductResponse
	var total int64
	var err error

	if qCategory != "" {
		if qPrice == "max" {
			products, total, err = ec.ecommerceRepo.FilterProductByCategoryAndPriceMax(qCategory, products, offset, pageSize)
			if err != nil {
				return &productResponses, total, err
			}
		} else if qPrice == "min" {
			products, total, err = ec.ecommerceRepo.FilterProductByCategoryAndPriceMin(qCategory, products, offset, pageSize)
			if err != nil {
				return &productResponses, total, err
			}
		} else {
			products, total, err = ec.ecommerceRepo.FilterProductByCategory(qCategory, products, offset, pageSize)
			if err != nil {
				return &productResponses, total, err
			}
		}
	} else {
		if qPrice == "max" {
			products, total, err = ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMax(products, offset, pageSize)
			if err != nil {
				return &productResponses, total, err
			}
		} else if qPrice == "min" {
			products, total, err = ec.ecommerceRepo.FilterProductByAllCategoryAndPriceMin(products, offset, pageSize)
			if err != nil {
				return &productResponses, total, err
			}
		}
	}

	for _, product := range *products {
		productImages, err := ec.ecommerceRepo.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return &productResponses, total, err
		}

		var imageURL string
		for _, image := range productImages {
			imageURL = image.ProductImageUrl
			break
		}

		productResponse := ee.ProductResponse{
			Name:            product.Name,
			Price:           product.Price,
			Rating:          product.Rating,
			ProductImageUrl: imageURL,
		}

		productResponses = append(productResponses, productResponse)
	}

	return &productResponses, total, nil
}
