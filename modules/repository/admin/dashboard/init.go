package dashboard

import (
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetDashboard(filter string) (int64, int64, int64, int64, error)
}

type dashboardRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) DashboardRepo {
	return &dashboardRepo{
		db,
	}
}
