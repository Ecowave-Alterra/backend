package dashboard

import de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"

func (dc *dashboardUsecase) GetDashboard(filter string) (int64, int64, int64, *[]de.FavouriteProducts, error) {
	totalRevenue, totalOrder, totalUser, top3Order, err := dc.dashboardRepo.GetDashboard(filter)

	return totalRevenue, totalOrder, totalUser, top3Order, err
}
