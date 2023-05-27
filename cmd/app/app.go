package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	adminHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/admin"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	adminRepo "github.com/berrylradianh/ecowave-go/modules/repository/admin"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	adminUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/admin"
	productUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUsecase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	adminRepo := adminRepo.New(mysql.DB)
	adminUsecase := adminUsecase.New(adminRepo)
	adminHandler := adminHandler.New(adminUsecase)

	handler := common.Handler{
		AdminHundler:   adminHandler,
		ProductHandler: productHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
