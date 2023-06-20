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

func (rr *reviewRepo) GetProductByID(productId string, product *pe.Product) (pe.Product, error) {
	if err := rr.db.
		Preload("ProductCategory").
		Where("id = ?", productId).
		First(&product).Error; err != nil {
		return *product, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetProductByName(name string, product *[]pe.Product) ([]pe.Product, error) {
	if err := rr.db.Where("name LIKE ?", "%"+name+"%").Preload("ProductCategory").
		Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetAllProductByCategory(category string, product *[]pe.Product) ([]pe.Product, error) {
	if err := rr.db.Preload("ProductCategory").
		Where("product_category_id IN (SELECT id FROM product_categories WHERE category = ?)", category).Find(&product).Error; err != nil {
		return nil, err
	}

	return *product, nil
}

func (rr *reviewRepo) GetAllTransactionDetailsNoPagination(productID string, transactionDetails *[]te.TransactionDetail) ([]te.TransactionDetail, error) {
	if err := rr.db.Where("product_id = ?", productID).Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return *transactionDetails, nil
}

func (rr *reviewRepo) GetAllTransactionDetail(productID string, transactionDetails *[]te.TransactionDetail, offset, pageSize int) ([]te.TransactionDetail, int64, error) {
	var count int64
	if err := rr.db.Model(&te.TransactionDetail{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := rr.db.Where("product_id = ?", productID).Offset(offset).Limit(pageSize).Find(&transactionDetails).Error; err != nil {
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

func (rr *reviewRepo) GetAllReviewByID(reviewID string, review *re.Review) (re.Review, error) {
	if err := rr.db.Where("id = ?", reviewID).Find(&review).Error; err != nil {
		return *review, err
	}

	return *review, nil
}
