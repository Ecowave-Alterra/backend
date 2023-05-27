package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	informationRepo "github.com/berrylradianh/ecowave-go/modules/repository/information"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	informationUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/information"
	productUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUsecase := productUsecase.New(productRepo)
	productHandler := productHandler.New(productUsecase)

	informationRepo := informationRepo.New(mysql.DB)
	informationUsecase := informationUsecase.New(informationRepo)
	informationHandler := informationHandler.New(informationUsecase)

	handler := common.Handler{
		ProductHandler:     productHandler,
		InformationHandler: informationHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
