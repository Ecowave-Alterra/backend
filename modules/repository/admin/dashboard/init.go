package information

import (
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetDashboard(filter string) error
}

type dashboardRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) DashboardRepo {
	return &dashboardRepo{
		db,
	}
}
