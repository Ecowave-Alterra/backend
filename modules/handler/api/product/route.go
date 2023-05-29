package product

import (
	// echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (productHandler *ProductHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	productGroup := e.Group("/products")
	productGroup.POST("/", productHandler.CreateProduct)
	productGroup.GET("/", productHandler.GetAllProduct)
	productGroup.GET("/:id", productHandler.GetProductByID)
	productGroup.PUT("/:id", productHandler.UpdateProduct)
	productGroup.DELETE("/:id", productHandler.DeleteProduct)
	// productGroup.POST("/import-csv", productHandler.ImportProductFromCSV)
	productGroup.GET("/search", productHandler.SearchProduct)
	productGroup.GET("/filter", productHandler.FilterProductByStatus)

}
