package transaction

import (
	"time"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
	"github.com/labstack/echo/v4"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) error {

	transactionDetail := transaction.TransactionDetails
	var productCost float64

	for _, cost := range transactionDetail {
		productCost += cost.SubTotalPrice
	}

	transaction.StatusTransaction = "Belum Bayar"
	transaction.TotalProductPrice = productCost

	transaction.TotalPrice = (transaction.TotalProductPrice + transaction.TotalShippingPrice) - (transaction.Point + transaction.Discount)

	err := tu.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}
func (tu *transactionUsecase) GetPoint(id uint) (interface{}, error) {

	res, err := tu.transactionRepo.GetPoint(id)
	if err != nil {
		return 0, err
	}

	if res == 0 {
		return "Maaf, Kamu tidak punya point", nil
	}

	return res, nil
}
func (tu *transactionUsecase) GetVoucherUser(id uint, offset int, pageSize int) (interface{}, int64, error) {

	res, count, err := tu.transactionRepo.GetVoucherUser(id, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	if res == nil {
		return "Kamu tidak mempunyai voucher", 0, nil
	}

	return res, count, nil
}

func (tu *transactionUsecase) DetailVoucher(id uint) (interface{}, error) {

	res, err := tu.transactionRepo.DetailVoucher(id)
	if err != nil {
		return nil, err
	}
	res.ID = 0
	if res.ID == 0 {
		return nil, echo.NewHTTPError(404, "Belum ada detail voucher")
	}

	detailVoucher := ev.VoucherUserResponse{
		Type:            res.VoucherType.Type,
		EndDate:         res.EndDate,
		PhotoUrl:        res.VoucherType.PhotoURL,
		MinimumPurchase: res.MinimumPurchase,
		MaximumDiscount: res.MaximumDiscount,
		DiscountPercent: res.DiscountPercent,
	}

	return detailVoucher, nil
}

func (tu *transactionUsecase) ClaimVoucher(idUser uint, idVoucher uint, shipCost float64, productCost float64) (float64, error) {

	var diskon float64

	res, err := tu.transactionRepo.ClaimVoucher(idVoucher)
	if err != nil {
		return 0, err
	}

	// validasi limit voucher user
	userClaim, err := tu.transactionRepo.CountVoucherUser(idUser, idVoucher)
	if err != nil {
		return 0, err
	}
	if res.ClaimableCount > userClaim {
		return 0, echo.NewHTTPError(400, "User telah melebihi batas penggunaan voucher")
	}
	// validasi limit voucher
	if res.MaxClaimLimit <= 0 {
		return 0, echo.NewHTTPError(400, "Voucher telah melebihi batas penggunaan")
	}
	// validasi tanggal
	now := time.Now()
	date := res.EndDate.Before(now)
	if !date {
		return 0, echo.NewHTTPError(400, "Telah melewati batas tanggal penggunaan voucher")
	}
	// validasi minimal belanjaan
	if res.MinimumPurchase > productCost {
		return 0, echo.NewHTTPError(400, "Total pembelian kurang untuk menggunakan voucher ini")
	}

	if res.VoucherType.Type == "Gratis Ongkir" {
		diskon = (shipCost * res.DiscountPercent) / 100
	}

	if res.VoucherType.Type == "Diskon Belanja" {
		diskon = (productCost * res.DiscountPercent) / 100

		if diskon > res.MaximumDiscount {
			diskon = res.MaximumDiscount
		}
	}

	return diskon, nil
}
