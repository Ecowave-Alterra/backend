package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	authHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/auth"
	informationHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/information"
	authRepo "github.com/berrylradianh/ecowave-go/modules/repository/auth"
	informationRepo "github.com/berrylradianh/ecowave-go/modules/repository/information"
	authUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/auth"
	informationUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/information"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	authRepo := authRepo.New(mysql.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	informationRepo := informationRepo.New(mysql.DB)
	informationUsecase := informationUsecase.New(informationRepo)
	informationHandler := informationHandler.New(informationUsecase)

	handler := common.Handler{
		AuthHandler:        authHandler,
		InformationHandler: informationHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
