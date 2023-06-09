package transaction

import (
	"errors"
	"log"
	"time"

	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ev "github.com/berrylradianh/ecowave-go/modules/entity/voucher"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (interface{}, error) {

	transactionDetail := transaction.TransactionDetails
	var productCost float64

	for _, cost := range transactionDetail {
		productCost += cost.SubTotalPrice
	}

	transaction.StatusTransaction = "Belum Bayar"
	transaction.TotalProductPrice = productCost

	var diskon float64

	transaction.TotalPrice = (transaction.TotalProductPrice + transaction.TotalShippingPrice) - (transaction.Point + diskon)
	log.Println(transaction)

	res, err := tu.transactionRepo.CreateTransaction(transaction)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (tu *transactionUsecase) GetPoint(id uint) (uint, error) {

	res, err := tu.transactionRepo.GetPoint(id)
	if err != nil {
		return 0, err
	}

	if res == 0 {
		return 0, errors.New("Maaf, Kamu tidak punya point")
	}

	return res, nil
}
func (tu *transactionUsecase) GetVoucherUser(id uint) (interface{}, error) {

	res, err := tu.transactionRepo.GetVoucherUser(id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return "Kamu tidak mempunyai voucher", nil
	}

	return res, nil
}
func (tu *transactionUsecase) DetailVoucher(id uint) (interface{}, error) {

	res, err := tu.transactionRepo.DetailVoucher(id)
	if err != nil {
		return nil, err
	}
	if res.ID == 0 {
		return "Belum ada detail voucher", nil
	}

	detailVoucher := ev.DetailVoucherResponse{
		Type:            res.VoucherType.Type,
		EndDate:         res.EndDate,
		PhotoUrl:        res.VoucherType.PhotoURL,
		MinimumPurchase: res.MinimumPurchase,
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
		return 0, errors.New("User telah melebihi batas penggunaan voucher")
	}
	// validasi limit voucher
	if res.MaxClaimLimit <= 0 {
		return 0, errors.New("Sudah melebihi batas penggunaan voucher")
	}
	// validasi tanggal
	now := time.Now()
	date := res.EndDate.Before(now)
	if !date {
		return 0, errors.New("Telah melewati batas tanggal penggunaan voucher")
	}
	// validasi minimal belanjaan
	if res.MinimumPurchase > productCost {
		return 0, errors.New("Total pembelian kurang untuk menggunakan voucher ini")
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
