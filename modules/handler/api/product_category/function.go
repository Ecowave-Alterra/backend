package productcategory

import (
	"fmt"
	"net/http"
	"strconv"

	pct "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (pch *ProductCategoryHandler) CreateProductCategory(c echo.Context) error {
	var productCategory pct.ProductCategory

	if err := c.Bind(&productCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := c.Validate(productCategory); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErr {
				if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required", e.Field())
				}
			}

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
			})
		}
	}

	exist, err := pch.productCategoryUsecase.CreateProductCategory(&productCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if exist {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "product category already exists",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new product category",
	})
}

func (pch *ProductCategoryHandler) UpdateProductCategory(c echo.Context) error {
	var productCategory pct.ProductCategory

	if err := c.Bind(&productCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if err := c.Validate(productCategory); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErr {
				if e.Tag() == "required" {
					message = fmt.Sprintf("%s is required", e.Field())
				}
			}

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
			})
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	exist, err := pch.productCategoryUsecase.UpdateProductCategory(&productCategory, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if exist {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "product category already exists",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product category by id",
	})
}

func (pch *ProductCategoryHandler) DeleteProductCategory(c echo.Context) error {
	var productCategory pct.ProductCategory

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := pch.productCategoryUsecase.DeleteProductCategory(&productCategory, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete product category by id",
	})
}

func (pch *ProductCategoryHandler) GetAllProductCategory(c echo.Context) error {
	var productCategory []pct.ProductCategory

	if err := pch.productCategoryUsecase.GetAllProductCategory(&productCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all product category",
		"data":    &productCategory,
	})
}

func (pch *ProductCategoryHandler) SearchingProductCategoyByName(c echo.Context) error {
	var productCategory []pct.ProductCategory

	name := c.QueryParam("name")

	available, err := pch.productCategoryUsecase.SearchingProductCategoryByName(&productCategory, name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}

	if !available {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "product category not available",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product category by name",
		"data":    &productCategory,
	})
}
