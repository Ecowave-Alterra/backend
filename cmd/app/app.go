package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	productCategoryHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product_category"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	productCategoryRepo "github.com/berrylradianh/ecowave-go/modules/repository/product_category"
	productUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	productCategoryUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product_category"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUsecase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	productCategoryRepo := productCategoryRepo.New(mysql.DB)
	productCategoryUsecase := productCategoryUsecase.New(productCategoryRepo)
	productCategoryHandler := productCategoryHandler.New(productCategoryUsecase)

	handler := common.Handler{
		ProductHandler:         productHandler,
		ProductCategoryHandler: productCategoryHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
