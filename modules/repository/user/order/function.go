package order

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (or *orderRepo) GetOrder(id string, idUser uint) ([]et.Transaction, error) {
	var transaction []et.Transaction

	err := or.db.Preload("TransactionDetails").Where("status_transaction = ? AND user_id = ?", id, idUser).Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
func (or *orderRepo) OrderDetail(id uint) (et.Transaction, error) {
	var transaction et.Transaction

	err := or.db.Preload("TransactionDetails").Where("id = ?", id).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (or *orderRepo) GetNameProductandImageUrl(id uint) (string, string, error) {
	var product ep.Product
	var pImg ep.ProductImage

	err := or.db.Select("name").Where("id = ?", id).First(&product).Error

	if err != nil {
		return "", "", err
	}
	err = or.db.Select("product_image_url").Where("product_id = ?", id).First(&pImg).Error
	if err != nil {
		return "", "", err
	}

	return product.Name, pImg.ProductImageUrl, nil
}

func (or *orderRepo) GetPromoName(id uint) (string, error) {
	var promo ev.Voucher

	err := or.db.Where("id = ?", id).First(&promo).Error

	if err != nil {
		return "", err
	}

	return promo.VoucherType.Type, nil
}
