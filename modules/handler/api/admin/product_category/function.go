package productcategory

import (
	"math"
	"net/http"
	"reflect"
	"strconv"

	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (pch *ProductCategoryHandler) GetAllProductCategory(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	productCategories, total, err := pch.productCategoryUsecase.GetAllProductCategory(offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	if productCategories == nil || len(*productCategories) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Belum ada list kategori",
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
		"ProductCategories": productCategories,
		"Page":              page,
		"TotalPage":         totalPages,
		"Status":            http.StatusOK,
	})
}

func (pch *ProductCategoryHandler) CreateProductCategory(c echo.Context) error {
	var productCategory pe.ProductCategory

	if err := c.Bind(&productCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusBadRequest,
		})
	}

	available, err := pch.productCategoryUsecase.CreateProductCategory(&productCategory)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	if available {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Kategori sudah ada",
			"Status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menambahkan kategori",
		"Status":  http.StatusCreated,
	})
}

func (pch *ProductCategoryHandler) UpdateProductCategory(c echo.Context) error {
	var productCategory pe.ProductCategory

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "ID harus berupa angka",
			"Status":  http.StatusBadRequest,
		})
	}

	_, err = pch.productCategoryUsecase.GetProductCategoryById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": err.Error(),
			"Status":  http.StatusNotFound,
		})
	}

	if err := c.Bind(&productCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusBadRequest,
		})
	}

	if reflect.DeepEqual(productCategory, pe.ProductCategory{}) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Masukkan kategori",
			"Status":  http.StatusBadRequest,
		})
	}

	available, err := pch.productCategoryUsecase.UpdateProductCategory(&productCategory, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	if available {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Kategori sudah ada",
			"Status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil mengubah kategori",
		"Status":  http.StatusOK,
	})
}

func (pch *ProductCategoryHandler) DeleteProductCategory(c echo.Context) error {
	var productCategory pe.ProductCategory

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "ID harus berupa angka",
			"Status":  http.StatusBadRequest,
		})
	}

	_, err = pch.productCategoryUsecase.GetProductCategoryById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusNotFound,
		})
	}

	if err := pch.productCategoryUsecase.DeleteProductCategory(&productCategory, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menghapus kategori",
		"Status":  http.StatusOK,
	})
}

func (pch *ProductCategoryHandler) SearchingProductCategoyByName(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	name := c.QueryParam("name")

	validParams := map[string]bool{"name": true, "page": true}
	for param := range c.QueryParams() {
		if !validParams[param] {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Masukkan parameter dengan benar",
				"Status":  http.StatusBadRequest,
			})
		}
	}

	productCategory, total, err := pch.productCategoryUsecase.SearchingProductCategoryByName(name, offset, pageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	if len(*productCategory) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Kategori yang anda cari tidak ditemukan",
			"Status":  http.StatusNotFound,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"ProductCategory": &productCategory,
			"Page":            page,
			"TotalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
			"Status":          http.StatusOK,
		})
	}
}
