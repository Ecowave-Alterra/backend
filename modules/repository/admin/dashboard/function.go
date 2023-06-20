package dashboard

import (
	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (dr *dashboardRepo) GetDashboard(filter string) (int64, int64, int64, *[]de.FavouriteProducts, *[]de.MonthlyRevenue, *[]de.WeeklyRevenue, error) {
	var totalRevenue int64
	var totalOrder int64
	var totalUser int64
	var monthlyRevenue *[]de.MonthlyRevenue
	var weeklyRevenue *[]de.WeeklyRevenue
	var top3Order *[]de.FavouriteProducts
	// var top3Review int64

	err := dr.db.Model(&te.Transaction{}).Select("COALESCE(SUM(total_price), 0) as total_income").
		Where("transactions.canceled_reason = ''").
		Row().Scan(&totalRevenue)
	if err != nil {
		return 0, 0, 0, nil, nil, nil, err
	}

	err = dr.db.Model(&te.Transaction{}).Count(&totalOrder).Error
	if err != nil {
		return 0, 0, 0, nil, nil, nil, err
	}

	err = dr.db.Model(&ue.User{}).Where("role_id = ?", 2).Count(&totalUser).Error
	if err != nil {
		return 0, 0, 0, nil, nil, nil, err
	}

	// query := "SELECT p.Name, SUM(td.Qty) AS TotalQty
	// FROM transactions AS t
	// JOIN transaction_details AS td ON t.id = td.transaction_id
	// JOIN products AS p ON p.id = td.product_id
	// WHERE t.canceled_reason IS NULL
	// GROUP BY p.name
	// ORDER BY TotalQty DESC
	// LIMIT 3"

	err = dr.db.Model(&te.Transaction{}).
		Select("products.name AS Name, SUM(transaction_details.qty) AS TotalOrders").
		Joins("JOIN transaction_details ON transactions.id = transaction_details.transaction_id").
		Joins("JOIN products ON products.product_id = transaction_details.product_id").
		Where("transactions.canceled_reason = ''").
		Group("products.name").
		Order("TotalOrders DESC").
		Limit(3).Scan(&top3Order).Error
	if err != nil {
		return 0, 0, 0, nil, nil, nil, err
	}

	// query := "SELECT CASE MONTH(transactions.created_at)
	// WHEN 1 THEN 'January'
	// WHEN 2 THEN 'February'
	// WHEN 3 THEN 'March'
	//  WHEN 4 THEN 'April'
	// WHEN 5 THEN 'May'
	//  WHEN 6 THEN 'June'
	// WHEN 7 THEN 'July'
	// WHEN 8 THEN 'August'
	// WHEN 9 THEN 'September'
	// WHEN 10 THEN 'October'
	// WHEN 11 THEN 'November'
	// WHEN 12 THEN 'December' END AS Month,
	// SUM(transactions.total_price) AS Revenue FROM transactions
	// WHERE YEAR(transactions.created_at) = YEAR(CURDATE()) AND transactions.canceled_reason IS NULL
	// GROUP BY MONTH(transactions.created_at), transactions.created_at
	// ORDER BY MONTH(transactions.created_at)"
	err = dr.db.Model(&te.Transaction{}).
		Select("CASE MONTH(transactions.created_at) " +
			"WHEN 1 THEN 'January' " +
			"WHEN 2 THEN 'February' " +
			"WHEN 3 THEN 'March' " +
			"WHEN 4 THEN 'April' " +
			"WHEN 5 THEN 'May' " +
			"WHEN 6 THEN 'June' " +
			"WHEN 7 THEN 'July' " +
			"WHEN 8 THEN 'August' " +
			"WHEN 9 THEN 'September' " +
			"WHEN 10 THEN 'October' " +
			"WHEN 11 THEN 'November' " +
			"WHEN 12 THEN 'December' " +
			"END AS Month, SUM(transactions.total_price) AS Revenue").
		Where("YEAR(transactions.created_at) = YEAR(CURDATE())").
		Where("transactions.canceled_reason = ''").
		Group("MONTH(transactions.created_at), transactions.created_at").
		Order("MONTH(transactions.created_at)").
		Scan(&monthlyRevenue).Error
	if err != nil {
		return 0, 0, 0, nil, nil, nil, err
	}

	// query := "SELECT DAYNAME(created_at) AS day, SUM(total_price) AS revenue
	// FROM transactions WHERE YEARWEEK(created_at) = YEARWEEK(CURDATE())
	// GROUP BY DAYNAME(created_at), DAYOFWEEK(created_at) ORDER BY DAYOFWEEK(created_at)"
	err = dr.db.Model(&te.Transaction{}).
		Select("DAYNAME(created_at) AS day, SUM(total_price) AS revenue").
		Where("YEARWEEK(created_at) = YEARWEEK(CURDATE())").
		Where("transactions.canceled_reason = ''").
		Group("DAYNAME(created_at), DAYOFWEEK(created_at)").
		Order("DAYOFWEEK(created_at)").
		Scan(&weeklyRevenue).Error
	if err != nil {
		return 0, 0, 0, nil, nil, nil, err
	}

	return totalRevenue, totalOrder, totalUser, top3Order, monthlyRevenue, weeklyRevenue, nil
}
