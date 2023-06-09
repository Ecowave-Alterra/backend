package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	mysql "github.com/berrylradianh/ecowave-go/database/mysql"
	ecommerceHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/ecommerce"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	ecommerceRepo "github.com/berrylradianh/ecowave-go/modules/repository/ecommerce"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	ecommerceUseCase "github.com/berrylradianh/ecowave-go/modules/usecase/ecommerce"
	productUseCase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUseCase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	ecommerceRepo := ecommerceRepo.New(mysql.DB)
	ecommerceUsecase := ecommerceUseCase.New(ecommerceRepo)
	ecommerceHandler := ecommerceHandler.New(ecommerceUsecase)

	handler := common.Handler{
		ProductHandler:   productHandler,
		EcommerceHandler: ecommerceHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
