package admin

import ac "github.com/berrylradianh/ecowave-go/modules/usecase/admin"

type AdminHandler struct {
	adminUsecase ac.AdminUsecase
}

func New(adminUsecase ac.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUsecase,
	}
}
