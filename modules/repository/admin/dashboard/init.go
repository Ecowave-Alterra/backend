package dashboard

import (
	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetDashboard() (int64, int64, int64, *[]de.FavouriteProducts, *[]de.TopReviews, error)
	GetYearlyRevenue() (*[]de.ChartResponse, error)
	GetMonthlyRevenue() (*[]de.ChartResponse, error)
	GetWeeklyRevenue() (*[]de.ChartResponse, error)
}

type dashboardRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) DashboardRepo {
	return &dashboardRepo{
		db,
	}
}
