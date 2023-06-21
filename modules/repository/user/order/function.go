package order

import (
	eo "github.com/berrylradianh/ecowave-go/modules/entity/order"
	ep "github.com/berrylradianh/ecowave-go/modules/entity/product"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	eu "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (or *orderRepo) GetOrder(filter string, idUser uint, offset int, pageSize int) (interface{}, int64, error) {
	var transaction []et.Transaction
	var order []eo.Order
	var OrderDetail []eo.OrderDetail
	var address eu.UserAddress
	var count int64
	err := or.db.Preload("TransactionDetails").Where("status_transaction = ? AND user_id = ?", filter, idUser).Find(&transaction).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = or.db.Preload("TransactionDetails").Where("status_transaction = ? AND user_id = ?", filter, idUser).Find(&transaction).Offset(offset).Limit(pageSize).Error
	if err != nil {
		return nil, 0, err
	}

	for _, val := range transaction {
		for _, td := range val.TransactionDetails {
			var pImg ep.ProductImage
			err = or.db.Select("product_image_url").Where("product_id = ?", td.ProductId).First(&pImg).Error
			if err != nil {
				return nil, 0, err
			}

			od := eo.OrderDetail{
				ProductId:       td.ProductId,
				ProductName:     td.ProductName,
				Qty:             td.Qty,
				SubTotalPrice:   td.SubTotalPrice,
				ProductImageUrl: pImg.ProductImageUrl,
				RatingProductId: td.RatingProductId,
			}
			OrderDetail = append(OrderDetail, od)
		}

		err = or.db.Where("id = ?", val.AddressId).First(&address).Error
		if err != nil {
			return nil, 0, err
		}

		ord := eo.Order{
			TransactionId:      val.TransactionId,
			CreatedAt:          val.CreatedAt,
			UpdatedAt:          val.UpdatedAt,
			AddressId:          val.AddressId,
			StatusTransaction:  val.StatusTransaction,
			ReceiptNumber:      val.ReceiptNumber,
			TotalProductPrice:  val.TotalProductPrice,
			TotalShippingPrice: val.TotalShippingPrice,
			Point:              val.Point,
			PaymentMethod:      val.PaymentMethod,
			PaymentStatus:      val.PaymentStatus,
			ExpeditionName:     val.ExpeditionName,
			ExpeditionStatus:   val.ExpeditionStatus,
			VoucherId:          val.VoucherId,
			Discount:           val.Discount,
			TotalPrice:         val.TotalPrice,
			CanceledReason:     val.CanceledReason,
			EstimationDay:      val.EstimationDay,
			PaymentUrl:         val.PaymentUrl,
			ExpeditionRating:   val.ExpeditionRating,
			OrderDetail:        OrderDetail,
			Address:            address,
		}
		order = append(order, ord)
	}

	return order, count, nil
}

func (or *orderRepo) GetStatusOrder(id string) (string, error) {

	var transaction et.Transaction
	err := or.db.Select("status_transaction").Where("transaction_id = ?", id).First(&transaction).Error
	if err != nil {
		return "", err
	}

	return transaction.StatusTransaction, nil
}
func (or *orderRepo) ConfirmOrder(id string) error {

	err := or.db.Model(&et.Transaction{}).Where("transaction_id = ?", id).Update("status_transaction", "Selesai").Error

	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepo) CancelOrder(co eo.CanceledOrder) error {

	err := or.db.Model(&et.Transaction{}).Where("transaction_id = ?", co.TransactionId).Updates(et.Transaction{StatusTransaction: "Dibatalkan", CanceledReason: co.CanceledReason}).Error

	if err != nil {
		return err
	}

	var transaction et.Transaction
	err = or.db.Where("transaction_id = ?", co.TransactionId).First(&transaction).Error
	if err != nil {
		return err
	}

	//update stock
	for _, val := range transaction.TransactionDetails {
		var product ep.Product
		err := or.db.Select("stock").Where("product_id = ?", val.ProductId).First(&product).Error
		if err != nil {
			return err
		}

		stock := product.Stock + val.Qty

		err = or.db.Model(&ep.Product{}).Where("product_id = ?", val.ProductId).Update("stock", stock).Error
		if err != nil {
			return err
		}
	}

	return nil
}
