package dashboard

import (
	dr "github.com/berrylradianh/ecowave-go/modules/repository/admin/dashboard"
)

type DashboardUsecase interface {
	GetDashboard(filter string) (int64, int64, int64, int64, error)
}

type dashboardUsecase struct {
	dashboardRepo dr.DashboardRepo
}

func New(dashboardRepo dr.DashboardRepo) *dashboardUsecase {
	return &dashboardUsecase{
		dashboardRepo,
	}
}
