package review

import (
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (rr *reviewRepo) CreateReview(review *er.Review) error {
	if err := rr.db.Create(&review).Error; err != nil {
		return err
	}

	return nil
}

func (rr *reviewRepo) CreateReviewDetail(reviewDetail *er.ReviewDetail) error {
	if err := rr.db.Create(&reviewDetail).Error; err != nil {
		return err
	}

	return nil
}

func (rr *reviewRepo) GetIdReview(idTransaction int) (int, error) {
	var review *er.Review

	if err := rr.db.Where("transaction_id = ?", idTransaction).Find(&review).Error; err != nil {
		return 0, err
	}

	return int(review.ID), nil
}

func (rr *reviewRepo) GetIdTransaction(transactionId string) (int, error) {
	var transaction *et.Transaction

	if err := rr.db.Where("transaction_id = ?", transactionId).Find(&transaction).Error; err != nil {
		return 0, err
	}

	return int(transaction.ID), nil
}

func (rr *reviewRepo) CountTransactionDetail(transactionId string) (int, error) {
	// var transaction *et.Transaction
	// var transactionDetail *et.TransactionDetail
	// subQuery := rr.db.Model(&transaction).Where("transaction_id = ?", transactionId).Select("id")
	// result := rr.db.Model(&transactionDetail).Where("transaction_id = ?", subQuery)
	// if result.Error != nil {
	// 	return 0, result.Error
	// }
	// count := result.RowsAffected

	var count int

	if err := rr.db.Raw("SELECT COUNT(td.product_id) FROM transaction_details td WHERE td.transaction_id = (SELECT id FROM transactions WHERE transaction_id = ?)", transactionId).Scan(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (rr *reviewRepo) GetPoint(id int) (int, error) {
	var userDetail *eu.UserDetail

	if err := rr.db.Where("id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}

	return int(userDetail.Point), nil
}

func (rr *reviewRepo) UpdatePoint(id int, point int) error {
	var userDetail *eu.UserDetail

	err := rr.db.Model(userDetail).Where("id = ?", id).Update("point", point).Error
	if err != nil {
		return err
	}

	return nil
}
