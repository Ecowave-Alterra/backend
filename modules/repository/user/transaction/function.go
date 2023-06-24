package transaction

import (
	"errors"

	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (tr *transactionRepo) GetUserById(id uint) (*ue.User, error) {
	user := &ue.User{}
	err := tr.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}

func (tr *transactionRepo) CreateTransaction(transaction *et.Transaction) error {

	err := tr.db.Create(&transaction).Error
	if err != nil {
		return err
	}

	//update stock
	for _, val := range transaction.TransactionDetails {
		var product ep.Product
		err := tr.db.Select("stock").Where("product_id = ?", val.ProductId).First(&product).Error
		if err != nil {
			return err
		}

		stock := product.Stock - val.Qty
		if stock == 0 {
			err = tr.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Updates(ep.Product{Stock: stock, Status: "habis"}).Error
		} else {
			err = tr.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Update("stock", stock).Error
		}

		if err != nil {
			return err
		}
	}
	return nil
}
func (tr *transactionRepo) UpdateTransaction(updateData et.Transaction) error {

	result := tr.db.Model(&et.Transaction{}).Where("transaction_id = ?", updateData.TransactionId).Updates(&updateData)

	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected < 1 {
		return errors.New("err")
	}
	return nil
}

func (tr *transactionRepo) GetPoint(id uint) (uint, error) {
	var userDetail eu.UserDetail

	if err := tr.db.Where("id = ?", id).First(&userDetail).Error; err != nil {
		return 0, err
	}
	point := userDetail.Point

	return point, nil

}
func (tr *transactionRepo) GetPaymentStatus(id string) (string, error) {
	var transaction et.Transaction

	if err := tr.db.Where("transaction_id = ?", id).First(&transaction).Error; err != nil {
		return "", err
	}
	status := transaction.PaymentStatus

	return status, nil

}

func (tr *transactionRepo) GetStock(id string) (uint, error) {
	var product ep.Product

	if err := tr.db.Where("product_id = ?", id).First(&product).Error; err != nil {
		return 0, err
	}
	stock := product.Stock

	return stock, nil

}

func (tr *transactionRepo) GetVoucherUser(id uint, offset int, pageSize int) ([]ev.Voucher, int64, error) {
	type IdVoucher struct {
		Id int
	}
	var idVoucher []IdVoucher
	var voucher []ev.Voucher
	var count int64

	subquery := tr.db.Model(&ev.Voucher{}).Select(`vouchers.id, voucher_types.type, vouchers.end_date, voucher_types.photo_url,
	vouchers.minimum_purchase,vouchers.max_claim_limit, count(*) user_claim`).Joins("left join transactions on transactions.voucher_id = vouchers.id").Joins("left join voucher_types on voucher_types.id = transactions.voucher_id").Where("transactions.user_id = ?", id).Group("vouchers.id")

	err := tr.db.Select(`id`).Table("(?) as sub", subquery).Where("user_claim > max_claim_limit").Scan(&idVoucher).Error
	if err != nil {
		return nil, 0, err
	}

	err = tr.db.Not(idVoucher).Find(&voucher).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = tr.db.Preload("VoucherType").Offset(offset).Limit(pageSize).Not(idVoucher).Find(&voucher).Error
	if err != nil {
		return voucher, 0, err
	}

	return voucher, count, nil

}
func (tr *transactionRepo) CountVoucherUser(idUser uint, idVoucher uint) (uint, error) {
	var count int64

	err := tr.db.Model(&et.Transaction{}).Where("user_id = ? AND voucher_id = ?", idUser, idVoucher).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint(count), nil

}
