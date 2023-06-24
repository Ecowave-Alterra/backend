package review

import (
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
)

func (rr *reviewRepo) GetAllProducts(offset, pageSize int) ([]re.GetAllReviewResponse, int64, error) {
	var reviews []re.GetAllReviewResponse
	var count int64

	// sql := SELECT p.product_id AS ProductID, p.name AS Name, pc.category AS Category, COUNT(rp.rating) AS ReviewQty
	// FROM products p
	// LEFT JOIN product_categories pc ON pc.id = p.product_category_id
	// LEFT JOIN transaction_details td ON td.product_id = p.product_id
	// LEFT JOIN rating_products rp ON rp.transaction_detail_id = td.id
	// GROUP BY ProductID, Name, Category;

	if err := rr.db.Table("products AS p").
		Select("p.product_id AS ProductID, p.name AS Name, pc.category AS Category, COUNT(rp.rating) AS ReviewQty").
		Joins("LEFT JOIN product_categories pc ON pc.id = p.product_category_id").
		Joins("LEFT JOIN transaction_details td ON td.product_id = p.product_id").
		Joins("LEFT JOIN rating_products rp ON rp.transaction_detail_id = td.id").
		Group("p.product_id, p.name, pc.category").
		Offset(offset).Limit(pageSize).
		Count(&count).
		Error; err != nil {
		return nil, 0, err
	}

	if err := rr.db.Table("products AS p").
		Select("p.product_id AS ProductID, p.name AS Name, pc.category AS Category, COUNT(rp.rating) AS ReviewQty").
		Joins("LEFT JOIN product_categories pc ON pc.id = p.product_category_id").
		Joins("LEFT JOIN transaction_details td ON td.product_id = p.product_id").
		Joins("LEFT JOIN rating_products rp ON rp.transaction_detail_id = td.id").
		Group("p.product_id, p.name, pc.category").
		Offset(offset).Limit(pageSize).
		Scan(&reviews).
		Error; err != nil {
		return nil, 0, err
	}

	return reviews, count, nil
}

func (rr *reviewRepo) SearchProduct(search string, offset, pageSize int) ([]re.GetAllReviewResponse, int64, error) {
	var reviews []re.GetAllReviewResponse
	var count int64

	if err := rr.db.Table("products AS p").
		Select("p.product_id AS ProductID, p.name AS Name, pc.category AS Category, COUNT(rp.rating) AS ReviewQty").
		Joins("LEFT JOIN product_categories pc ON pc.id = p.product_category_id").
		Joins("LEFT JOIN transaction_details td ON td.product_id = p.product_id").
		Joins("LEFT JOIN rating_products rp ON rp.transaction_detail_id = td.id").
		Where("p.name LIKE ? OR p.product_id LIKE ? OR pc.category LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%").
		Group("p.product_id, p.name, pc.category").
		Offset(offset).Limit(pageSize).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := rr.db.Table("products AS p").
		Select("p.product_id AS ProductID, p.name AS Name, pc.category AS Category, COUNT(rp.rating) AS ReviewQty").
		Joins("LEFT JOIN product_categories pc ON pc.id = p.product_category_id").
		Joins("LEFT JOIN transaction_details td ON td.product_id = p.product_id").
		Joins("LEFT JOIN rating_products rp ON rp.transaction_detail_id = td.id").
		Where("p.name LIKE ? OR p.product_id LIKE ? OR pc.category LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%").
		Group("p.product_id, p.name, pc.category").
		Offset(offset).Limit(pageSize).
		Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, count, nil
}

func (rr *reviewRepo) GetProductReviewById(productId string, offset, pageSize int) ([]re.ReviewResponse, int64, error) {
	var review []re.ReviewResponse
	var count int64

	// sql := SELECT t.transaction_id AS TransactionID,t.receipt_number AS ReceiptNumber,ud.name AS Name,
	// ud.profile_photo AS ProfilePhoto, p.name AS ProductName, pc.category AS ProductCategory,
	// rp.comment AS CommentUser, rp.comment_admin AS CommentAdmin, rp.photo_url AS PhotoUrl,
	// rp.video_url AS VideoUrl, CASE WHEN (t.expedition_rating * rp.rating) / 2 > 5 THEN 5 ELSE (t.expedition_rating * rp.rating) / 2 END AS AvgRating,
	// t.expedition_rating AS ExpeditionRating, rp.rating AS ProductRating
	// FROM transactions t
	// JOIN users u ON t.user_id = u.id
	// JOIN user_details ud ON ud.user_id = u.id
	// JOIN transaction_details td ON td.transaction_id = t.id
	// JOIN products p ON p.product_id = td.product_id
	// JOIN product_categories pc ON pc.id = p.product_category_id
	// JOIN rating_products rp ON rp.transaction_detail_id = td.id
	// WHERE t.status_transaction = "selesai" AND p.product_id LIKE "%a3325f33-e01a-4e40-9ca7-5d84c4337094%";

	if err := rr.db.Table("transactions AS t").
		Select("t.transaction_id AS TransactionID, t.receipt_number AS ReceiptNumber, ud.name AS Name, ud.profile_photo AS ProfilePhoto, p.name AS ProductName, pc.category AS ProductCategory, rp.comment AS CommentUser, rp.comment_admin AS CommentAdmin, rp.photo_url AS PhotoUrl, rp.video_url AS VideoUrl, CASE WHEN (t.expedition_rating * rp.rating) / 2 > 5 THEN 5 ELSE (t.expedition_rating * rp.rating) / 2 END AS AvgRating, t.expedition_rating AS ExpeditionRating, rp.rating AS ProductRating").
		Joins("JOIN users u ON t.user_id = u.id").
		Joins("JOIN user_details ud ON ud.user_id = u.id").
		Joins("JOIN transaction_details td ON td.transaction_id = t.id").
		Joins("JOIN products p ON p.product_id = td.product_id").
		Joins("JOIN product_categories pc ON pc.id = p.product_category_id").
		Joins("JOIN rating_products rp ON rp.transaction_detail_id = td.id").
		Where("t.status_transaction = ? AND p.product_id LIKE ?", "selesai", "%"+productId+"%").
		Offset(offset).Limit(pageSize).
		Count(&count).
		Error; err != nil {
		return review, 0, err
	}

	if err := rr.db.Table("transactions AS t").
		Select("t.transaction_id AS TransactionID, t.receipt_number AS ReceiptNumber, ud.name AS Name, ud.profile_photo AS ProfilePhoto, p.name AS ProductName, pc.category AS ProductCategory, rp.comment AS CommentUser, rp.comment_admin AS CommentAdmin, rp.photo_url AS PhotoUrl, rp.video_url AS VideoUrl, CASE WHEN (t.expedition_rating * rp.rating) / 2 > 5 THEN 5 ELSE (t.expedition_rating * rp.rating) / 2 END AS AvgRating, t.expedition_rating AS ExpeditionRating, rp.rating AS ProductRating").
		Joins("JOIN users u ON t.user_id = u.id").
		Joins("JOIN user_details ud ON ud.user_id = u.id").
		Joins("JOIN transaction_details td ON td.transaction_id = t.id").
		Joins("JOIN products p ON p.product_id = td.product_id").
		Joins("JOIN product_categories pc ON pc.id = p.product_category_id").
		Joins("JOIN rating_products rp ON rp.transaction_detail_id = td.id").
		Where("t.status_transaction = ? AND p.product_id LIKE ?", "selesai", "%"+productId+"%").
		Offset(offset).Limit(pageSize).
		Scan(&review).
		Error; err != nil {
		return review, 0, err
	}

	return review, count, nil
}
