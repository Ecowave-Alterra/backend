package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	productCategoryHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/admin/product_category"
	productCategoryRepo "github.com/berrylradianh/ecowave-go/modules/repository/admin/product_category"
	productCategoryUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/admin/product_category"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productCategoryRepo := productCategoryRepo.New(mysql.DB)
	productCategoryUsecase := productCategoryUsecase.New(productCategoryRepo)
	productCategoryHandler := productCategoryHandler.New(productCategoryUsecase)

	handler := common.Handler{
		ProductCategoryHandler: productCategoryHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
