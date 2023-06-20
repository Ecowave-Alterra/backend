package dashboard

import de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"

func (dc *dashboardUsecase) GetDashboard(filter string) (int64, int64, int64, *[]de.FavouriteProducts, *[]de.MonthlyRevenue, *[]de.WeeklyRevenue, error) {
	totalRevenue, totalOrder, totalUser, top3Order, monthlyRevenue, weeklyRevenue, err := dc.dashboardRepo.GetDashboard(filter)

	return totalRevenue, totalOrder, totalUser, top3Order, monthlyRevenue, weeklyRevenue, err
}
