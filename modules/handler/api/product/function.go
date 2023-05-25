package product

import (
	"net/http"

	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	"github.com/labstack/echo/v4"
)

func (productHandler *ProductHandler) GetAllProducts() echo.HandlerFunc {
	return func(e echo.Context) error {
		var products *[]ep.Product

		products, err := productHandler.productUsecase.GetAllProducts()
		if err != nil {
			return e.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Success Get All Products",
			"products": products,
		})
	}
}
