package order

import (
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"github.com/labstack/echo/v4"
)

func (or *orderRepo) GetOrder(id string, idUser uint, offset int, pageSize int) ([]et.Transaction, int64, error) {
	var transaction []et.Transaction
	var count int64

	err := or.db.Preload("TransactionDetails").Where("status_transaction = ? AND user_id = ?", id, idUser).Count(&count).Error
	if err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	err = or.db.Offset(offset).Limit(pageSize).Preload("TransactionDetails").Where("status_transaction = ? AND user_id = ?", id, idUser).Find(&transaction).Error
	if err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return transaction, count, nil
}
func (or *orderRepo) OrderDetail(id uint) (et.Transaction, error) {
	var transaction et.Transaction

	err := or.db.Preload("TransactionDetails").Where("id = ?", id).Find(&transaction).Error
	if err != nil {
		return transaction, echo.NewHTTPError(404, err)
	}

	return transaction, nil
}

func (or *orderRepo) GetNameProductandImageUrl(id uint) (string, string, error) {
	var product ep.Product
	var pImg ep.ProductImage

	err := or.db.Select("name").Where("id = ?", id).First(&product).Error

	if err != nil {
		return "", "", echo.NewHTTPError(404, err)
	}
	err = or.db.Select("product_image_url").Where("product_id = ?", id).First(&pImg).Error
	if err != nil {
		return "", "", echo.NewHTTPError(404, err)
	}

	return product.Name, pImg.ProductImageUrl, nil
}

func (or *orderRepo) GetPromoName(id uint) (string, error) {
	var promo ev.Voucher

	err := or.db.Where("id = ?", id).First(&promo).Error

	if err != nil {
		return "", echo.NewHTTPError(404, err)
	}

	return promo.VoucherType.Type, nil
}
func (or *orderRepo) GetStatusOrder(id uint) (string, error) {

	var transaction et.Transaction
	err := or.db.Select("status_transaction").Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return "", echo.NewHTTPError(404, err)
	}

	return transaction.StatusTransaction, nil
}
func (or *orderRepo) ConfirmOrder(id uint) error {

	err := or.db.Model(&et.Transaction{}).Where("id = ?", id).Update("status_transaction", "Selesai").Error

	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (or *orderRepo) CancelOrder(id uint, canceledReason string) error {

	err := or.db.Model(&et.Transaction{}).Where("id = ?", id).Updates(et.Transaction{StatusTransaction: "Dibatalkan", CanceledReason: canceledReason}).Error

	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	var transaction et.Transaction
	err = or.db.Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return echo.NewHTTPError(404, err)
	}

	//update stock
	for _, val := range transaction.TransactionDetails {
		var product ep.Product
		err := or.db.Select("stock").Where("product_id = ?", val.ProductId).First(&product).Error
		if err != nil {
			return echo.NewHTTPError(404, err)
		}

		stock := product.Stock + val.Qty

		err = or.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Update("stock", stock).Error
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
	}

	return nil
}
