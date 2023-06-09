package routes

import (
	"github.com/berrylradianh/ecowave-go/common"
	m "github.com/berrylradianh/ecowave-go/middleware/log"

	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = &vld.CustomValidator{Validator: validator.New()}
	m.LogMiddleware(e)
	e.Use(middleware.Recover())

	// productGroup := e.Group("/products")
	// productGroup.POST("/", handler.ProductHandler.CreateProduct)
	// productGroup.GET("/", handler.ProductHandler.GetAllProduct)
	// productGroup.GET("/:id", handler.ProductHandler.GetProductByID)
	// productGroup.PUT("/:id", handler.ProductHandler.UpdateProduct)
	// productGroup.DELETE("/:id", handler.ProductHandler.DeleteProduct)
	// productGroup.POST("/import-csv", handler.ProductHandler.ImportProductFromCSV)
	// productGroup.GET("/search", handler.ProductHandler.SearchProduct)
	// productGroup.GET("/filter", handler.ProductHandler.FilterProductByStatus)

	handler.ProductHandler.RegisterRoutes(e)
	handler.EcommerceHandler.RegisterRoutes(e)

	return e
}
