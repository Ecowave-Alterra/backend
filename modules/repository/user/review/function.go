package review

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (rr *reviewRepo) CountTransactionDetail(transactionId string) (int, error) {
	var count int

	if err := rr.db.Raw("SELECT COUNT(td.product_id) FROM transaction_details td WHERE td.transaction_id = (SELECT id FROM transactions WHERE transaction_id = ?)", transactionId).Scan(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (rr *reviewRepo) GetIdTransaction(transactionId string) (int, error) {
	var transaction *et.Transaction

	if err := rr.db.Where("transaction_id = ?", transactionId).Find(&transaction).Error; err != nil {
		return 0, err
	}

	return int(transaction.ID), nil
}

func (rr *reviewRepo) GetProductId(transactionId string) ([]string, error) {
	var productId []string

	if err := rr.db.Raw("SELECT td.product_id FROM transaction_details td WHERE td.transaction_id = (SELECT id FROM transactions WHERE transaction_id = ?)", transactionId).Scan(&productId).Error; err != nil {
		return nil, err
	}

	return productId, nil
}

func (rr *reviewRepo) GetIdTransactionDetail(idTransaction int, productId string) (int, error) {
	var id int

	if err := rr.db.Raw("SELECT id FROM transaction_details WHERE transaction_id = ? AND product_id = ?", idTransaction, productId).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (rr *reviewRepo) CreateRatingProduct(ratingProduct *er.RatingProduct) error {
	if err := rr.db.Create(&ratingProduct).Error; err != nil {
		return err
	}

	return nil
}

func (rr *reviewRepo) UpdateExpeditionRating(ratingExpedition float32, transactionId string) error {
	var transaction *et.Transaction

	err := rr.db.Model(transaction).Where("transaction_id = ?", transactionId).Update("expedition_rating", ratingExpedition).Error
	if err != nil {
		return err
	}

	return nil
}

func (rr *reviewRepo) GetPoint(id int) (int, error) {
	var userDetail *eu.UserDetail

	if err := rr.db.Where("id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}

	return int(userDetail.Point), nil
}

func (rr *reviewRepo) UpdatePoint(idUser int, point int) error {
	var userDetail *eu.UserDetail

	err := rr.db.Model(userDetail).Where("id = ?", idUser).Update("point", point).Error
	if err != nil {
		return err
	}

	return nil
}
