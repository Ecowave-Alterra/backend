package dashboard

import (
	dc "github.com/berrylradianh/ecowave-go/modules/usecase/admin/dashboard"
)

type DashboardHandler struct {
	dashboardUsecase dc.DashboardUsecase
}

func New(dashboardUsecase dc.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{
		dashboardUsecase,
	}
}
