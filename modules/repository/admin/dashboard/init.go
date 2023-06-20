package dashboard

import (
	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetDashboard(filter string) (int64, int64, int64, *[]de.FavouriteProducts, *[]de.MonthlyRevenue, *[]de.WeeklyRevenue, *[]de.YearlyRevenue, error)
}

type dashboardRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) DashboardRepo {
	return &dashboardRepo{
		db,
	}
}
