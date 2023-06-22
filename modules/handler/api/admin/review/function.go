package review

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) GetAllReview(c echo.Context) error {
	var products []pe.Product

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	products, total, err := rh.reviewUsecase.GetAllProducts(&products, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data produk",
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list produk",
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

	count := 0
	var transactionDetails []te.TransactionDetail
	var reviewResponses []re.GetAllReviewResponse
	for _, product := range products {
		transactionDetails, err := rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ID), &transactionDetails)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data transaksi detail produk",
			})
		}
		for _, td := range transactionDetails {
			if fmt.Sprint(td.RatingProductId) != "" {
				count++
			}
		}
		reviewResponse := re.GetAllReviewResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Category:  product.ProductCategory.Category,
			ReviewQty: uint(count),
		}

		reviewResponses = append(reviewResponses, reviewResponse)
		count = 0
	}

	if reviewResponses == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Belum ada list ulasan",
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

func (rh *ReviewHandler) GetReviewByProductID(c echo.Context) error {
	productID := c.Param("id")

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	var transactionDetails []te.TransactionDetail
	transactionDetails, total, err := rh.reviewUsecase.GetAllTransactionDetail(fmt.Sprint(productID), &transactionDetails, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data transaksi detail produk",
		})
	}

	if len(transactionDetails) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list detail transaksi",
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

	var review re.Review
	var reviewResponses []re.ReviewResponse
	var transaction te.Transaction
	var product pe.Product
	var user ue.User
	for _, td := range transactionDetails {
		if fmt.Sprint(td.RatingProductId) != "" {
			review, err = rh.reviewUsecase.GetAllReviewByID(fmt.Sprint(td.RatingProductId), &review)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data ulasan produk",
					"Status":  http.StatusInternalServerError,
				})
			}

			transaction, err := rh.reviewUsecase.GetTransactionByID(td.TransactionId, &transaction)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data transaksi",
					"Status":  http.StatusInternalServerError,
				})
			}

			product, err := rh.reviewUsecase.GetProductByID(fmt.Sprint(td.ProductId), &product)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data produk",
				})
			}

			user, err := rh.reviewUsecase.GetUserByID(fmt.Sprint(transaction.UserId), &user)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data user",
					"Status":  http.StatusInternalServerError,
				})
			}

			reviewResponse := re.ReviewResponse{
				TransactionID:    fmt.Sprint(td.TransactionId),
				Name:             user.Username,
				ProfilePhoto:     user.UserDetail.ProfilePhoto,
				ProductName:      product.Name,
				ProductCategory:  product.ProductCategory.Category,
				CommentUser:      review.Comment,
				CommentAdmin:     review.CommentAdmin,
				PhotoUrl:         review.PhotoUrl,
				VideoUrl:         review.VideoUrl,
				AvgRating:        product.Rating,
				ExpeditionRating: transaction.ExpeditionRating,
				ProductRating:    review.Rating,
			}

			reviewResponses = append(reviewResponses, reviewResponse)
		}
	}

	if reviewResponses == nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data ulasan produk",
			"Status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Reviews":   reviewResponses,
		"Page":      page,
		"TotalPage": totalPages,
		"Status":    http.StatusOK,
	})
}

func (rh *ReviewHandler) SearchReview(c echo.Context) error {
	param := c.QueryParam("param")

	switch param {
	case "id":
		productID := c.QueryParam("id")

		var product pe.Product
		product, err := rh.reviewUsecase.GetProductByID(fmt.Sprint(productID), &product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		var transactionDetails []te.TransactionDetail
		transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(productID), &transactionDetails)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data transaksi detail produk",
				"Status":  http.StatusInternalServerError,
			})
		}

		count := 0
		for _, td := range transactionDetails {
			if fmt.Sprint(td.RatingProductId) != "" {
				count++
			}
		}

		if count == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		reviewResponse := re.GetAllReviewResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Category:  product.ProductCategory.Category,
			ReviewQty: uint(count),
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Reviews": reviewResponse,
			"Status":  http.StatusOK,
		})
	case "name":
		productName := c.QueryParam("name")
		var products []pe.Product

		products, err := rh.reviewUsecase.GetProductByName(productName, &products)
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
			transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ID), &transactionDetails)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data transaksi detail produk",
					"Status":  http.StatusInternalServerError,
				})
			}
			for _, td := range transactionDetails {
				if fmt.Sprint(td.RatingProductId) != "" {
					count++
				}
			}

			reviewResponse := re.GetAllReviewResponse{
				ProductID: product.ID,
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

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Reviews": reviewResponses,
			"Status":  http.StatusOK,
		})
	case "category":
		productCategory := c.QueryParam("category")
		var products []pe.Product

		products, err := rh.reviewUsecase.GetAllProductByCategory(productCategory, &products)
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
			transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetailsNoPagination(fmt.Sprint(product.ID), &transactionDetails)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"Message": "Gagal mengambil data transaksi detail produk",
					"Status":  http.StatusInternalServerError,
				})
			}
			for _, td := range transactionDetails {
				if fmt.Sprint(td.RatingProductId) != "" {
					count++
				}
			}

			reviewResponse := re.GetAllReviewResponse{
				ProductID: product.ID,
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

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Reviews": reviewResponses,
			"Status":  http.StatusOK,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"Message": "Invalid search parameter",
		"Status":  http.StatusBadRequest,
	})
}
