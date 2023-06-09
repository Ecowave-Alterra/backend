package ecommerce

import (
	"fmt"
	"net/http"

	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (eh *EcommerceHandler) GetProductEcommerce(c echo.Context) error {
	var products []ep.Product

	products, err := eh.ecommerceUseCase.GetAllProduct(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to get product datas",
			"Error":   err,
		})
	}

	var productResponses []ee.ProductResponse
	var productImage ep.ProductImage

	for _, product := range products {
		productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to get product images",
				"Error":   err,
			})
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Successfully get product datas",
		"Products": productResponses,
	})
}

func (eh *EcommerceHandler) GetProductDetailEcommerce(c echo.Context) error {
	var product ep.Product
	productID := c.Param("id")

	product, err := eh.ecommerceUseCase.GetProductByID(productID, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to get product datas",
			"Error":   err,
		})
	}

	var productImage ep.ProductImage
	productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to get product images",
			"Error":   err,
		})
	}

	var imageURLs []string
	for _, image := range productImages {
		imageURLs = append(imageURLs, image.ProductImageUrl)
	}

	productDetailResponse := ee.ProductDetailResponse{
		Name:            product.Name,
		Category:        product.ProductCategory.Category,
		Stock:           product.Stock,
		Price:           product.Price,
		Status:          product.Status,
		Rating:          product.Rating,
		Description:     product.Description,
		ProductImageUrl: imageURLs,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Successfully get detail product datas",
		"Products": productDetailResponse,
	})
}

func (eh *EcommerceHandler) FilterProductByCategoryAndPrice(c echo.Context) error {
	var products []ep.Product
	var productResponses []ee.ProductResponse

	qCategory := c.QueryParam("category")
	qPrice := c.QueryParam("price")

	if qCategory != "" {
		if qPrice == "max" {
			products, err := eh.ecommerceUseCase.FilterProductByCategoryAndPriceMax(qCategory, &products)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product",
					"Error":   err,
				})
			}

			var productImage ep.ProductImage
			for _, product := range products {
				productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Failed to get product images",
						"Error":   err,
					})
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
		} else if qPrice == "min" {
			products, err := eh.ecommerceUseCase.FilterProductByCategoryAndPriceMin(qCategory, &products)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product",
					"Error":   err,
				})
			}

			var productImage ep.ProductImage
			for _, product := range products {
				productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Failed to get product images",
						"Error":   err,
					})
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
		} else {
			products, err := eh.ecommerceUseCase.FilterProductByCategory(qCategory, &products)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product",
					"Error":   err,
				})
			}

			var productImage ep.ProductImage
			for _, product := range products {
				productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Failed to get product images",
						"Error":   err,
					})
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
		}
	} else {
		if qPrice == "max" {
			products, err := eh.ecommerceUseCase.FilterProductByAllCategoryAndPriceMax(&products)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product",
					"Error":   err,
				})
			}

			var productImage ep.ProductImage
			for _, product := range products {
				productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Failed to get product images",
						"Error":   err,
					})
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
		} else if qPrice == "min" {
			products, err := eh.ecommerceUseCase.FilterProductByAllCategoryAndPriceMin(&products)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Message": "Failed to get product",
					"Error":   err,
				})
			}

			var productImage ep.ProductImage
			for _, product := range products {
				productImages, err := eh.ecommerceUseCase.GetProductImageURLById(fmt.Sprint(product.ID), &productImage)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, echo.Map{
						"Message": "Failed to get product images",
						"Error":   err,
					})
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
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Successfully get detail product datas",
		"Products": productResponses,
	})
}
