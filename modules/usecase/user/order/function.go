package order

import (
	o "github.com/berrylradianh/ecowave-go/modules/entity/order"
	"github.com/labstack/echo/v4"
)

func (oc *orderUsecase) GetOrder(filter string, idUser uint, offset int, pageSize int) (interface{}, int64, error) {

	res, count, err := oc.orderRepo.GetOrder(filter, idUser, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	if res == nil {
		return "Belum ada pesanan", 0, nil
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
			return nil, 0, err
		}
		order := o.OrderResponse{
			ProductImageUrl: imageUrl,
			PaymentStatus:   val.StatusTransaction,
			ProductName:     nameProduct,
			ProductQty:      productQty,
			ProductPrice:    float64(productPrice),
			TotalQty:        totalProduct,
			Total:           val.TotalPrice,
		}

		OrderResponse = append(OrderResponse, order)

	}

	return OrderResponse, count, nil
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
		PaymentMethod:     res.PaymentMethod,
		ExpeditionName:    res.ExpeditionName,
		ExpeditionStatus:  res.ExpeditionStatus,
		VoucherId:         res.VoucherId,
		AddressId:         res.AddressId,
		StatusTransaction: res.StatusTransaction,
		ShippingCost:      res.TotalShippingPrice,
		ProductCost:       res.TotalProductPrice,
		Point:             res.Point,
		TotalPrice:        res.TotalPrice,
		TotalProduct:      totalProduct,
		PromoName:         promoName,
		ReceiptNumber:     res.ReceiptNumber,
		Discount:          res.Discount,
		OrderDetails:      OrderDetailResponse,
	}

	return order, nil
}
func (oc *orderUsecase) ConfirmOrder(id uint) error {

	statusTransaction, err := oc.orderRepo.GetStatusOrder(id)
	if err != nil {
		return err
	}

	if statusTransaction != "Dikirim" {
		return echo.NewHTTPError(400, "Tidak bisa mengonfirmasi pesanan sebelum barang di terima")
	}

	err = oc.orderRepo.ConfirmOrder(id)
	if err != nil {
		return err
	}

	return nil
}
func (oc *orderUsecase) CancelOrder(id uint, canceledReason string) error {

	statusTransaction, err := oc.orderRepo.GetStatusOrder(id)
	if err != nil {
		return err
	}

	if statusTransaction != "Belum Bayar" {
		return echo.NewHTTPError(400, "Tidak bisa membatalkan pesanan")
	}

	err = oc.orderRepo.CancelOrder(id, canceledReason)
	if err != nil {
		return err
	}

	return nil
}
