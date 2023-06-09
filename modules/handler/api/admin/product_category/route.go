package productcategory

import "github.com/labstack/echo/v4"

func (pch *ProductCategoryHandler) RegisterRoutes(e *echo.Echo) {
	// jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	productCategoryGroup := e.Group("/admin/products/category")
	productCategoryGroup.POST("", pch.CreateProductCategory)
	productCategoryGroup.PUT("/:id", pch.UpdateProductCategory)
	productCategoryGroup.DELETE("/:id", pch.DeleteProductCategory)
	productCategoryGroup.GET("", pch.GetAllProductCategory)
	productCategoryGroup.GET("/", pch.SearchingProductCategoyByName)
}
