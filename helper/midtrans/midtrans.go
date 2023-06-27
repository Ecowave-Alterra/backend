package midtrans

import (
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateMidtransUrl(transaction *et.Transaction) (string, error) {

	var itemDetails *[]midtrans.ItemDetails
	var iDetails []midtrans.ItemDetails
	for _, val := range transaction.TransactionDetails {
		item := midtrans.ItemDetails{
			ID:    val.ProductId,
			Name:  val.ProductName,
			Qty:   int32(val.Qty),
			Price: int64(val.SubTotalPrice) / int64(val.Qty),
		}
		iDetails = append(iDetails, item)

	}
	shipping := midtrans.ItemDetails{
		ID:    transaction.ExpeditionName,
		Name:  "ongkir",
		Qty:   1,
		Price: int64(transaction.TotalShippingPrice),
	}
	iDetails = append(iDetails, shipping)

	if transaction.VoucherId != 0 {
		discount := midtrans.ItemDetails{
			ID:    "discount",
			Name:  "discount",
			Qty:   1,
			Price: -int64(transaction.Discount),
		}
		iDetails = append(iDetails, discount)
	}

	if transaction.Point != 0 {
		point := midtrans.ItemDetails{
			ID:    "point",
			Name:  "point",
			Qty:   1,
			Price: -int64(transaction.Point),
		}
		iDetails = append(iDetails, point)
	}

	itemDetails = &iDetails

	var s = snap.Client{}
	s.New("SB-Mid-server-2p-XclqW-K668VvZ6BSNfecI", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.TransactionId,
			GrossAmt: int64(transaction.TotalPrice),
		},
		Items: itemDetails,
	}

	snapResp, midtransErr := s.CreateTransactionUrl(req)
	if midtransErr != nil {
		return "", echo.NewHTTPError(500, "Gagal membuat transaksi")
	}

	return snapResp, nil
}
