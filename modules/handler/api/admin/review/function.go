package review

import (
	"fmt"
	"log"
	"net/http"

	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) GetAllReview(c echo.Context) error {
	var products []pe.Product
	products, err := rh.reviewUsecase.GetAllProducts(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data produk",
		})
	}

	count := 0
	var transactionDetails []te.TransactionDetail
	var reviewResponses []re.GetAllReviewResponse
	for _, product := range products {
		transactionDetails, err := rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(product.ID), &transactionDetails)
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
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Belum ada list ulasan",
			"Status":  http.StatusOK,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Berhasil mengambil data review produk",
		"Reviews": reviewResponses,
	})
}

func (rh *ReviewHandler) GetReviewByProductID(c echo.Context) error {
	productID := c.Param("id")

	var transactionDetails []te.TransactionDetail
	transactionDetails, err := rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(productID), &transactionDetails)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal mengambil data transaksi detail produk",
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

			log.Println(td.TransactionId)
			transaction, err := rh.reviewUsecase.GetTransactionByID(fmt.Sprint(td.TransactionId), &transaction)
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
		"Message": "Berhasil mengambil data review produk",
		"Reviews": reviewResponses,
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
			})
		}

		var transactionDetails []te.TransactionDetail
		transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(productID), &transactionDetails)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data transaksi detail produk",
			})
		}

		count := 0
		for _, td := range transactionDetails {
			if fmt.Sprint(td.RatingProductId) != "" {
				count++
			}
		}

		if count == 0 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusOK,
			})
		}

		reviewResponse := re.GetAllReviewResponse{
			ProductID: product.ID,
			Name:      product.Name,
			Category:  product.ProductCategory.Category,
			ReviewQty: uint(count),
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil mengambil data review produk",
			"Reviews": reviewResponse,
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
			transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(product.ID), &transactionDetails)
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusOK,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil mengambil data review produk",
			"Reviews": reviewResponses,
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
			transactionDetails, err = rh.reviewUsecase.GetAllTransactionDetails(fmt.Sprint(product.ID), &transactionDetails)
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
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Ulasan yang anda cari tidak ditemukan",
				"Status":  http.StatusOK,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil mengambil data review produk",
			"Reviews": reviewResponses,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"Message": "Invalid search parameter",
	})
}
