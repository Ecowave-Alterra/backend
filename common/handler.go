package common

import (
	userHandler "github.com/berrylradianh/ecowave-go/modules/handler/api/user"
)

type Handler struct {
	UserHandler *userHandler.UserHandler
}
