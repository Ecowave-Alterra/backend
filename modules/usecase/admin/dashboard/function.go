package dashboard

import (
	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"
)

func (dc *dashboardUsecase) GetDashboard() (int64, int64, int64, *[]de.FavouriteProducts, *[]de.TopReviews, error) {
	totalRevenue, totalOrder, totalUser, top3Order, top3Review, err := dc.dashboardRepo.GetDashboard()

	return totalRevenue, totalOrder, totalUser, top3Order, top3Review, err
}

func (dc *dashboardUsecase) GetYearlyRevenue() (*[]de.ChartResponse, error) {
	yearlyRevenue, err := dc.dashboardRepo.GetYearlyRevenue()

	return yearlyRevenue, err
}

func (dc *dashboardUsecase) GetMonthlyRevenue() (*[]de.ChartResponse, error) {
	monthlyRevenue, err := dc.dashboardRepo.GetMonthlyRevenue()

	return monthlyRevenue, err
}

func (dc *dashboardUsecase) GetWeeklyRevenue() (*[]de.ChartResponse, error) {
	weeklyRevenue, err := dc.dashboardRepo.GetWeeklyRevenue()

	return weeklyRevenue, err
}
