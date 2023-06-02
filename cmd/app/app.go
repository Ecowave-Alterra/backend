package app

import (
	"github.com/berrylradianh/ecowave-go/cmd/routes"
	"github.com/berrylradianh/ecowave-go/common"

	"github.com/berrylradianh/ecowave-go/database/mysql"
	userHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user"
	userRepo "github.com/berrylradianh/ecowave-go/modules/repository/user"
	userUsecase "github.com/berrylradianh/ecowave-go/modules/usecase/user"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	userRepo := userRepo.New(mysql.DB)
	userUsecase := userUsecase.New(userRepo)
	userHandler := userHandler.New(userUsecase)

	handler := common.Handler{
		UserHandler: userHandler,
	}

	router := routes.StartRoute(handler)
	return router
}
