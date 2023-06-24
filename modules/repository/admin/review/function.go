package review

import (
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	re "github.com/berrylradianh/ecowave-go/modules/entity/review"
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"github.com/labstack/echo/v4"
)

func (rr *reviewRepo) GetAllProducts(products *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	var count int64
	if err := rr.db.Model(&pe.Product{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := rr.db.Preload("ProductCategory").Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *products, count, nil
}

func (rr *reviewRepo) GetProductByIDNoPagination(productId string, product *pe.Product) (pe.Product, error) {
	if err := rr.db.
		Preload("ProductCategory").
		Where("product_id = ?", productId).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetProductByID(productId string, product *pe.Product, offset, pageSize int) (*pe.Product, int64, error) {
	var count int64
	if err := rr.db.Model(&pe.Product{}).Where("product_id LIKE ?", "%"+productId+"%").Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := rr.db.Preload("ProductCategory").Offset(offset).Limit(pageSize).Where("product_id LIKE ?", "%"+productId+"%").Find(&product).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return product, count, nil
}

func (rr *reviewRepo) GetProductByName(name string, product *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	var count int64
	if err := rr.db.Model(&pe.Product{}).Where("name LIKE ?", "%"+name+"%").Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := rr.db.Preload("ProductCategory").Offset(offset).Limit(pageSize).Where("name LIKE ?", "%"+name+"%").Find(&product).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *product, count, nil
}

func (rr *reviewRepo) GetAllProductByCategory(category string, product *[]pe.Product, offset, pageSize int) ([]pe.Product, int64, error) {
	var count int64
	if err := rr.db.Model(&pe.Product{}).Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := rr.db.Preload("ProductCategory").Offset(offset).Limit(pageSize).Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *product, count, nil
	// if err := rr.db.Preload("ProductCategory").
	// 	Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
	// 	return nil, err
	// }

	// return *product, nil
}

func (rr *reviewRepo) GetAllTransactionDetailsNoPagination(productID string, transactionDetails []te.TransactionDetail) ([]te.TransactionDetail, error) {
	if err := rr.db.Where("product_id = ?", productID).Preload("RatingProduct").Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return transactionDetails, nil
}

func (rr *reviewRepo) GetAllTransactionDetail(productID string, transactionDetails *[]te.TransactionDetail, offset, pageSize int) ([]te.TransactionDetail, int64, error) {
	var count int64
	if err := rr.db.Model(&te.TransactionDetail{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := rr.db.Where("product_id = ?", productID).Offset(offset).Limit(pageSize).Preload("RatingProduct").Find(&transactionDetails).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *transactionDetails, count, nil
}

func (rr *reviewRepo) GetTransactionByID(transactionID uint, transaction *te.Transaction) (te.Transaction, error) {
	if err := rr.db.Where("id = ?", transactionID).First(&transaction).Error; err != nil {
		return *transaction, err
	}

	return *transaction, nil
}

func (rr *reviewRepo) GetUserByID(userID string, user *ue.User) (ue.User, error) {
	if err := rr.db.Preload("UserDetail").Where("id = ?", userID).First(&user).Error; err != nil {
		return *user, err
	}

	return *user, nil
}

func (rr *reviewRepo) GetAllReviewByID(reviewID string, review *re.RatingProduct) (re.RatingProduct, error) {
	if err := rr.db.Where("id = ?", reviewID).Find(&review).Error; err != nil {
		return *review, err
	}

	return *review, nil
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
