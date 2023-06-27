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

	err := dr.db.Raw(" SELECT months.month_name AS Label,COALESCE(SUM(transactions.total_price), 0) AS Value FROM ( SELECT 1 AS month_number, 'January' AS month_name UNION SELECT 2, 'February' UNION SELECT 3, 'March' UNION SELECT 4, 'April' UNION SELECT 5, 'May' UNION SELECT 6, 'June' UNION SELECT 7, 'July' UNION SELECT 8, 'August' UNION SELECT 9, 'September' UNION SELECT 10, 'October' UNION SELECT 11, 'November' UNION SELECT 12, 'December') AS months LEFT JOIN (SELECT MONTH(created_at) AS month_number,SUM(total_price) AS total_price FROM transactions WHERE YEAR(created_at) = YEAR(CURDATE()) AND canceled_reason = '' AND deleted_at IS NULL GROUP BY month_number) AS transactions ON months.month_number = transactions.month_number GROUP BY months.month_number, months.month_name ORDER BY months.month_number").
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
