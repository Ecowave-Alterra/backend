package transaction

import (
	"log"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (tr *transactionRepo) CreateTransaction(transaction *et.Transaction) (interface{}, error) {
	log.Println(transaction)

	err := tr.db.Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (tr *transactionRepo) GetPoint(id uint) (uint, error) {
	var userDetail eu.UserDetail

	if err := tr.db.Where("id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}
	point := userDetail.Point

	return point, nil

}

func (tr *transactionRepo) GetVoucherUser(id uint) ([]ev.VoucherUserResponse, error) {
	var result []ev.VoucherUserResponse

	subquery := tr.db.Model(&ev.Voucher{}).Select(`vouchers.id, voucher_types.type, vouchers.end_date, voucher_types.photo_url,
	vouchers.minimum_purchase,vouchers.max_claim_limit, count(*) user_claim`).Joins("left join transactions on transactions.voucher_id = vouchers.id").Joins("left join voucher_types on voucher_types.id = transactions.voucher_id").Where("transactions.user_id = ?", id).Group("transactions.voucher_id")

	err := tr.db.Table("(?) as sub", subquery).Where("user_claim < max_claim_limit").Scan(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil

}
func (tr *transactionRepo) CountVoucherUser(idUser uint, idVoucher uint) (uint, error) {
	var count int64

	err := tr.db.Model(&et.Transaction{}).Where("user_id = ? AND voucher_id = ?", idUser, idVoucher).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint(count), nil

}
func (tr *transactionRepo) ClaimVoucher(id uint) (ev.Voucher, error) {
	var voucher ev.Voucher

	if err := tr.db.Where("id = ?", id).First(&voucher).Error; err != nil {
		return voucher, err
	}

	return voucher, nil

}
func (tr *transactionRepo) DetailVoucher(id uint) (ev.Voucher, error) {
	var voucher ev.Voucher

	if err := tr.db.Preload("VoucherType").Where("id = ?", id).First(&voucher).Error; err != nil {
		return voucher, err
	}

	return voucher, nil

}
