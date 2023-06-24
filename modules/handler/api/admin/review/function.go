package review

import (
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (rh *ReviewHandler) GetAllProducts(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	products, total, err := rh.reviewUsecase.GetAllProducts(offset, pageSize)
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

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list review",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Reviews":   products,
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

	search := c.QueryParam("search")
	validParams := map[string]bool{"search": true, "page": true}
	for param := range c.QueryParams() {
		if !validParams[param] {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Masukkan parameter dengan benar",
				"Status":  http.StatusBadRequest,
			})
		}
	}

	products, total, err := rh.reviewUsecase.SearchProduct(search, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": err,
			"Status":  http.StatusInternalServerError,
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Product yang anda cari tidak ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Reviews":   products,
		"Page":      page,
		"TotalPage": int(math.Ceil(float64(total) / float64(pageSize))),
		"Status":    http.StatusOK,
	})
}
