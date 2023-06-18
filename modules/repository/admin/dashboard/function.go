package dashboard

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (dr *dashboardRepo) GetDashboard(filter string) (int64, int64, int64, error) {
	var totalIncome int64
	var totalOrder int64
	var totalUser int64
	// var weeklyIncome int64
	// var top3Order int64
	// var top3Review int64

	err := dr.db.Model(&te.Transaction{}).Select("COALESCE(SUM(total_price), 0) as total_income").Row().Scan(&totalIncome)
	if err != nil {
		return 0, 0, 0, err
	}

	err = dr.db.Model(&te.Transaction{}).Count(&totalOrder).Error
	if err != nil {
		return 0, 0, 0, err
	}

	err = dr.db.Model(&ue.User{}).Where("role_id = ?", 2).Count(&totalUser).Error
	if err != nil {
		return 0, 0, 0, err
	}

	return totalIncome, totalOrder, totalUser, nil
}
