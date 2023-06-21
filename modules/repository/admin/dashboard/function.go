package dashboard

import (
	de "github.com/berrylradianh/ecowave-go/modules/entity/dashboard"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (dr *dashboardRepo) GetDashboard() (int64, int64, int64, *[]de.FavouriteProducts, *[]de.TopReviews, error) {
	var totalRevenue int64
	var totalOrder int64
	var totalUser int64
	var top3Order *[]de.FavouriteProducts
	var top3Review *[]de.TopReviews

	err := dr.db.Model(&te.Transaction{}).Select("COALESCE(SUM(total_price), 0) as total_income").
		Where("transactions.canceled_reason = ''").
		Row().Scan(&totalRevenue)
	if err != nil {
		return 0, 0, 0, nil, nil, err
	}

	err = dr.db.Model(&te.Transaction{}).Count(&totalOrder).Error
	if err != nil {
		return 0, 0, 0, nil, nil, err
	}

	err = dr.db.Model(&ue.User{}).Where("role_id = ?", 2).Count(&totalUser).Error
	if err != nil {
		return 0, 0, 0, nil, nil, err
	}

	err = dr.db.Model(&te.Transaction{}).
		Select("products.name AS Name, SUM(transaction_details.qty) AS TotalOrders").
		Joins("JOIN transaction_details ON transactions.id = transaction_details.transaction_id").
		Joins("JOIN products ON products.product_id = transaction_details.product_id").
		Where("transactions.canceled_reason = ''").
		Group("products.name").
		Order("TotalOrders DESC").
		Limit(3).Scan(&top3Order).Error
	if err != nil {
		return 0, 0, 0, nil, nil, err
	}

	err = dr.db.Model(&te.Transaction{}).
		Select("products.name AS Name, COUNT(transaction_details.qty) AS TotalReviews").
		Joins("JOIN transaction_details ON transactions.id = transaction_details.transaction_id").
		Joins("JOIN rating_products ON rating_products.transaction_detail_id = transaction_details.id").
		Joins("JOIN products ON products.product_id = transaction_details.product_id").
		Group("products.name").
		Order("TotalReviews DESC").
		Limit(3).Scan(&top3Review).Error
	if err != nil {
		return 0, 0, 0, nil, nil, err
	}

	return totalRevenue, totalOrder, totalUser, top3Order, top3Review, nil
}

func (dr *dashboardRepo) GetYearlyRevenue() (*[]de.ChartResponse, error) {
	var yearlyRevenue *[]de.ChartResponse

	err := dr.db.Model(&te.Transaction{}).
		Select("YEAR(created_at) AS Label, SUM(total_price) AS Value").
		Where("YEAR(created_at) BETWEEN YEAR(CURDATE()) - 7 AND YEAR(CURDATE())").
		Where("canceled_reason = ''").
		Group("YEAR(created_at)").
		Order("YEAR(created_at)").
		Scan(&yearlyRevenue).Error
	if err != nil {
		return nil, err
	}

	return yearlyRevenue, nil
}

func (dr *dashboardRepo) GetMonthlyRevenue() (*[]de.ChartResponse, error) {
	var monthlyRevenue *[]de.ChartResponse

	err := dr.db.Model(&te.Transaction{}).
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
			"END AS Label, SUM(transactions.total_price) AS Value").
		Where("YEAR(transactions.created_at) = YEAR(CURDATE())").
		Where("transactions.canceled_reason = ''").
		Group("MONTH(transactions.created_at), transactions.created_at").
		Order("MONTH(transactions.created_at)").
		Scan(&monthlyRevenue).Error
	if err != nil {
		return nil, err
	}
	return monthlyRevenue, nil
}

func (dr *dashboardRepo) GetWeeklyRevenue() (*[]de.ChartResponse, error) {
	var weeklyRevenue *[]de.ChartResponse

	err := dr.db.Model(&te.Transaction{}).
		Select("DAYNAME(created_at) AS Label, SUM(total_price) AS Value").
		Where("YEARWEEK(created_at) = YEARWEEK(CURDATE())").
		Where("transactions.canceled_reason = ''").
		Group("DAYNAME(created_at), DAYOFWEEK(created_at)").
		Order("DAYOFWEEK(created_at)").
		Scan(&weeklyRevenue).Error
	if err != nil {
		return nil, err
	}

	return weeklyRevenue, nil
}
