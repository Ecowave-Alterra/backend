package dashboard

func (dc *dashboardUsecase) GetDashboard(filter string) (int64, int64, int64, error) {
	totalIncome, totalOrder, totalUser, err := dc.dashboardRepo.GetDashboard(filter)

	return totalIncome, totalOrder, totalUser, err
}
