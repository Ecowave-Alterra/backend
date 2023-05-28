package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"
	mysql "github.com/berrylradianh/ecowave-go/database/mysql"
	productHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/product"
	productRepo "github.com/berrylradianh/ecowave-go/modules/repository/product"
	productUseCase "github.com/berrylradianh/ecowave-go/modules/usecase/product"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	productRepo := productRepo.New(mysql.DB)
	productUseCase := productUseCase.New(productRepo)
	productHandler := productHandler.New(productUseCase)

	handler := common.Handler{
		ProductHandler: productHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
