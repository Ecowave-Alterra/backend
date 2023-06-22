package dashboard

import (
	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"
	dr "github.com/berrylradianh/ecowave-go/modules/repository/admin/dashboard"
)

type DashboardUsecase interface {
	GetDashboard() (int64, int64, int64, *[]de.FavouriteProducts, *[]de.TopReviews, error)
	GetYearlyRevenue() (*[]de.ChartResponse, error)
	GetMonthlyRevenue() (*[]de.ChartResponse, error)
	GetWeeklyRevenue() (*[]de.ChartResponse, error)
}

type dashboardUsecase struct {
	dashboardRepo dr.DashboardRepo
}

func New(dashboardRepo dr.DashboardRepo) *dashboardUsecase {
	return &dashboardUsecase{
		dashboardRepo,
	}
}
