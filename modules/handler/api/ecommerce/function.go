package ecommerce

import (
	"math"
	"net/http"
	"strconv"

	ee "github.com/berrylradianh/ecowave-go/modules/entity/ecommerce"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (eh *EcommerceHandler) GetProductEcommerce(c echo.Context) error {
	var products *[]ep.Product
	var productResponses *[]ee.ProductResponse

	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	productResponses, total, err := eh.ecommerceUseCase.GetAllProduct(products, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to get product datas",
			"Error":   err,
		})
	}

	// code, msg := cs.CustomStatus(err.Error())
	// if err != nil {
	// 	return c.JSON(code, echo.Map{
	// 		"Status":  code,
	// 		"Message": msg,
	// 	})
	// }

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	if productResponses == nil || len(*productResponses) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Belum ada list produk",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Products":  productResponses,
		"Page":      page,
		"Status":    http.StatusOK,
		"TotalPage": totalPages,
	})
}

func (eh *EcommerceHandler) GetProductDetailEcommerce(c echo.Context) error {
	productID := c.Param("id")

	productDetailResponse, err := eh.ecommerceUseCase.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to get product datas",
			"Error":   err,
		})
	}

	// code, msg := cs.CustomStatus(err.Error())
	// if err != nil {
	// 	return c.JSON(code, echo.Map{
	// 		"Status":  code,
	// 		"Message": msg,
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Products": productDetailResponse,
		"Status":   http.StatusOK,
	})
}

func (eh *EcommerceHandler) FilterProductByCategoryAndPrice(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	qCategory := c.QueryParam("category")
	qPrice := c.QueryParam("price")

	productResponses, total, err := eh.ecommerceUseCase.FilterProductByCategoryAndPrice(qCategory, qPrice, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Failed to get product",
			"Error":   err,
		})
	}

	// code, msg := cs.CustomStatus(err.Error())
	// if err != nil {
	// 	return c.JSON(code, echo.Map{
	// 		"Status":  code,
	// 		"Message": msg,
	// 	})
	// }

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > totalPages {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Halaman Tidak Ditemukan",
			"Status":  http.StatusNotFound,
		})
	}

	if productResponses == nil || len(*productResponses) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Belum ada list produk",
			"Status":  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Products":  productResponses,
		"Page":      page,
		"Status":    http.StatusOK,
		"TotalPage": totalPages,
	})
}
