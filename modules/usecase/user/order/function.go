package order

import (
	o "github.com/berrylradianh/ecowave-go/modules/entity/order"
)

func (oc *orderUsecase) GetOrder(id string, idUser uint) (interface{}, error) {

	res, err := oc.orderRepo.GetOrder(id, idUser)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return "Belum ada pesanan", nil
	}

	var OrderResponse []o.OrderResponse
	var idProduct, totalProduct, productPrice, productQty uint

	for _, val := range res {
		for i, td := range val.TransactionDetails {
			if i == 0 {
				idProduct = td.ProductId
				productPrice = uint(td.SubTotalPrice)
				productQty = td.Qty
			}
			totalProduct += td.Qty
		}

		nameProduct, imageUrl, err := oc.orderRepo.GetNameProductandImageUrl(idProduct)
		if err != nil {
			return nil, err
		}
		order := o.OrderResponse{
			ProductImageUrl:      imageUrl,
			PaymentStatus:        val.StatusTransaction,
			ProductName:          nameProduct,
			ProductQty:           productQty,
			ProductPrice:         float64(productPrice),
			TransactionDetailQty: totalProduct,
			Total:                val.TotalPrice,
		}

		OrderResponse = append(OrderResponse, order)

	}

	return OrderResponse, nil
}

func (oc *orderUsecase) OrderDetail(id uint) (interface{}, error) {

	res, err := oc.orderRepo.OrderDetail(id)
	if err != nil {
		return nil, err
	}

	var idProduct, totalProduct uint
	var OrderDetailResponse []o.OrderDetailResponse

	for _, val := range res.TransactionDetails {
		idProduct = val.ProductId
		totalProduct += val.Qty

		nameProduct, imageUrl, err := oc.orderRepo.GetNameProductandImageUrl(idProduct)
		if err != nil {
			return nil, err
		}

		orderDetail := o.OrderDetailResponse{
			TransactionId:   val.TransactionId,
			ProductId:       val.ProductId,
			Qty:             val.Qty,
			SubTotalPrice:   val.SubTotalPrice,
			NameProduct:     nameProduct,
			ProductImageUrl: imageUrl,
		}

		OrderDetailResponse = append(OrderDetailResponse, orderDetail)

	}
	promoName, err := oc.orderRepo.GetPromoName(res.VoucherId)
	if err != nil {
		return nil, err
	}

	order := o.Order{
		UserId:        res.UserId,
		PaymentMethod: res.PaymentMethod,
		// ExpeditionId:       res.ExpeditionId,
		VoucherId:         res.VoucherId,
		AddressId:         res.AddressId,
		StatusTransaction: res.StatusTransaction,
		// ShippingCost:       res.ShippingCost,
		// ProductCost:        res.ProductCost,
		Point:              res.Point,
		TotalPrice:         res.TotalPrice,
		TotalProduct:       totalProduct,
		PromoName:          promoName,
		TransactionDetails: OrderDetailResponse,
	}

	return order, nil
}
