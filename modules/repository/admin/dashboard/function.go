package dashboard

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (dr *dashboardRepo) GetDashboard(filter string) (int64, int64, int64, int64, error) {
	var totalRevenue int64
	var totalOrder int64
	var totalUser int64
	// var weeklyIncome int64
	var top3Order int64
	// var top3Review int64

	err := dr.db.Model(&te.Transaction{}).Select("COALESCE(SUM(total_price), 0) as total_income").Row().Scan(&totalRevenue)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	err = dr.db.Model(&te.Transaction{}).Count(&totalOrder).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	err = dr.db.Model(&ue.User{}).Where("role_id = ?", 2).Count(&totalUser).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// query := "SELECT p.Name, SUM(td.Qty) AS TotalQty FROM transactions AS t JOIN transaction_details AS td ON t.id = td.transaction_id JOIN products AS p ON p.id = td.product_id WHERE t.canceled_reason IS NULL GROUP BY p.name ORDER BY TotalQty DESC LIMIT 3"
	err = dr.db.Model(&te.Transaction{}).
		Select("products.name AS Name, SUM(transaction_details.qty) AS TotalQty").
		Joins("JOIN transaction_details ON transactions.id = transaction_details.transaction_id").
		Joins("JOIN products ON products.id = transaction_details.product_id").
		Where("transactions.canceled_reason IS NULL").
		Group("products.name").
		Order("TotalQty DESC").
		Limit(3).Scan(&top3Order).Error
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return totalRevenue, totalOrder, totalUser, top3Order, nil
}
