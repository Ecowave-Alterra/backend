package review

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) GetAllProductReviews(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	reviews, total, err := rh.reviewUsecase.GetAllProductReviews(offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
			"Status":  http.StatusInternalServerError,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	if len(reviews) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list review",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Reviews":   reviews,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

// func (rh *ReviewHandler) GetAllReview(c echo.Context) error {
// 	var products []pe.Product

// 	pageParam := c.QueryParam("page")
// 	page, err := strconv.Atoi(pageParam)
// 	if err != nil || page < 1 {
// 		page = 1
// 	}

// 	pageSize := 10
// 	offset := (page - 1) * pageSize

// 	products, total, err := rh.reviewUsecase.GetAllProducts(&products, offset, pageSize)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"Message": "Gagal mengambil data produk",
// 		})
// 	}

// 	if len(products) == 0 {
// 		return c.JSON(http.StatusNotFound, echo.Map{
// 			"Message": "Belum ada list produk",
// 			"Status":  http.StatusNotFound,
// 		})
// 	}

// 	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
// 	if page > totalPages {
// 		return c.JSON(http.StatusNotFound, echo.Map{
// 			"Message": "Halaman Tidak Ditemukan",
// 			"Status":  http.StatusNotFound,
// 		})
// 	}

// 	count := 0
// 	var transactionDetails []te.TransactionDetail
// 	var reviewResponses []re.GetAllReviewResponse
// 	for _, product := range products {
// 		transactionDetails, err := rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ProductId), transactionDetails)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 				"Message": "Gagal mengambil data transaksi detail produk",
// 			})
// 		}
// 		for _, td := range transactionDetails {
// 			if td.RatingProduct.Rating != 0 {
// 				count++
// 			}
// 		}
// 		reviewResponse := re.GetAllReviewResponse{
// 			ProductID: product.ProductId,
// 			Name:      product.Name,
// 			Category:  product.ProductCategory.Category,
// 			ReviewQty: uint(count),
// 		}

// 		reviewResponses = append(reviewResponses, reviewResponse)
// 		count = 0
// 	}

// 	if reviewResponses == nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"Message": "Belum ada list ulasan",
// 			"Status":  http.StatusNotFound,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"Reviews":   reviewResponses,
// 		"Page":      page,
// 		"TotalPage": totalPages,
// 		"Status":    http.StatusOK,
// 	})
// }

func (rh *ReviewHandler) GetReviewByProductID(c echo.Context) error {
	productID := c.Param("id")

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	reviews, total, err := rh.reviewUsecase.GetProductReviewById(productID, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
			"Status":  http.StatusInternalServerError,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	if len(reviews) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list review",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Reviews":   reviews,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

func (rh *ReviewHandler) SearchReview(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	param := c.QueryParam("param")

	switch param {
	case "id":
		productID := c.QueryParam("id")

		var product *pe.Product
		product, total, err := rh.reviewUsecase.GetProductByID(fmt.Sprint(productID), product, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		var transactionDetails []te.TransactionDetail
		transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ProductId), transactionDetails)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data transaksi detail produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		count := 0
		for _, td := range transactionDetails {
			if fmt.Sprint(td.RatingProduct.ID) != "" {
				count++
			}
		}

		if count == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		reviewResponse := re.GetAllReviewResponse{
			ProductID: product.ProductId,
			Name:      product.Name,
			Category:  product.ProductCategory.Category,
			ReviewQty: uint(count),
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Reviews":   reviewResponse,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	case "name":
		productName := c.QueryParam("name")
		var products []pe.Product

		products, total, err := rh.reviewUsecase.GetProductByName(productName, &products, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		count := 0
		var transactionDetails []te.TransactionDetail
		var reviewResponses []re.GetAllReviewResponse
		for _, product := range products {
			transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ProductId), transactionDetails)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data transaksi detail produk",
					"Status":  http.StatusInternalServerError,
				})
			}
			for _, td := range transactionDetails {
				if fmt.Sprint(td.RatingProduct.ID) != "" {
					count++
				}
			}

			reviewResponse := re.GetAllReviewResponse{
				ProductID: product.ProductId,
				Name:      product.Name,
				Category:  product.ProductCategory.Category,
				ReviewQty: uint(count),
			}

			reviewResponses = append(reviewResponses, reviewResponse)

			count = 0
		}
		if reviewResponses == nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Reviews":   reviewResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	case "category":
		productCategory := c.QueryParam("category")
		var products []pe.Product

		products, total, err := rh.reviewUsecase.GetAllProductByCategory(productCategory, &products, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		count := 0
		var transactionDetails []te.TransactionDetail
		var reviewResponses []re.GetAllReviewResponse
		for _, product := range products {
			transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ProductId), transactionDetails)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data transaksi detail produk",
					"Status":  http.StatusInternalServerError,
				})
			}
			for _, td := range transactionDetails {
				if fmt.Sprint(td.RatingProduct.ID) != "" {
					count++
				}
			}

			reviewResponse := re.GetAllReviewResponse{
				ProductID: product.ProductId,
				Name:      product.Name,
				Category:  product.ProductCategory.Category,
				ReviewQty: uint(count),
			}

			reviewResponses = append(reviewResponses, reviewResponse)

			count = 0
		}
		if reviewResponses == nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Reviews":   reviewResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"Message": "Invalid search parameter",
		"Status":  http.StatusBadRequest,
	})
}
