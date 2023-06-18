package transaction

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/berrylradianh/ecowave-go/helper/hash"
	mdtrns "github.com/berrylradianh/ecowave-go/helper/midtrans"
	"github.com/berrylradianh/ecowave-go/helper/rajaongkir"
	em "github.com/berrylradianh/ecowave-go/modules/entity/midtrans"
	er "github.com/berrylradianh/ecowave-go/modules/entity/rajaongkir"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"

	"github.com/labstack/echo/v4"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (string, string, error) {
	var productCost float64

	for _, cost := range transaction.TransactionDetails {
		productCost += cost.SubTotalPrice
	}

	transId := "eco" + strconv.FormatUint(uint64(transaction.UserId), 10) + time.Now().UTC().Format("2006010215040105")
	transaction.TransactionId = transId
	transaction.StatusTransaction = "Belum Bayar"
	transaction.TotalProductPrice = productCost
	transaction.TotalPrice = (transaction.TotalProductPrice + transaction.TotalShippingPrice) - (transaction.Point + transaction.Discount)

	redirectUrl, err := mdtrns.CreateMidtransUrl(transaction)
	if err != nil {
		return "", "", err
	}
	transaction.PaymentUrl = redirectUrl

	err = tu.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return "", "", err
	}
	return redirectUrl, transId, nil
}
func (tu *transactionUsecase) MidtransNotifications(midtransRequest *em.MidtransRequest) error {

	Key := hash.Hash(midtransRequest.OrderId, midtransRequest.StatusCode, midtransRequest.GrossAmount)

	if Key != midtransRequest.SignatureKey {
		log.Println("sini?", Key)
		log.Println(midtransRequest.SignatureKey)
		return echo.NewHTTPError(400, "Invalid Transaction")
	}

	transaction := et.Transaction{
		TransactionId: midtransRequest.OrderId,
		PaymentStatus: midtransRequest.TransactionStatus,
	}
	if midtransRequest.TransactionStatus == "settlement" {
		transaction.StatusTransaction = "Dikemas"
	}
	err := tu.transactionRepo.UpdateTransaction(transaction)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Invalid Transaction")
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
func (tu *transactionUsecase) GetPaymentStatus(id string) (string, error) {

	res, err := tu.transactionRepo.GetPaymentStatus(id)
	if err != nil {
		return "", err
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

// func (tu *transactionUsecase) DetailVoucher(id uint) (interface{}, error) {

// 	res, err := tu.transactionRepo.DetailVoucher(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if res.ID == 0 {
// 		return nil, echo.NewHTTPError(404, "Belum ada detail voucher")
// 	}

// 	detailVoucher := ev.VoucherUserResponse{
// 		Id:              res.ID,
// 		Type:            res.VoucherType.Type,
// 		EndDate:         res.EndDate,
// 		PhotoUrl:        res.VoucherType.PhotoURL,
// 		MinimumPurchase: res.MinimumPurchase,
// 		MaximumDiscount: res.MaximumDiscount,
// 		DiscountPercent: res.DiscountPercent,
// 	}

// 	return detailVoucher, nil
// }

// func (tu *transactionUsecase) ClaimVoucher(idUser uint, idVoucher uint, shipCost float64, productCost float64) (float64, error) {

// 	var diskon float64

// 	res, err := tu.transactionRepo.ClaimVoucher(idVoucher)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// validasi limit voucher user
// 	userClaim, err := tu.transactionRepo.CountVoucherUser(idUser, idVoucher)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if res.ClaimableCount > userClaim {
// 		return 0, echo.NewHTTPError(400, "User telah melebihi batas penggunaan voucher")
// 	}
// 	// validasi limit voucher
// 	if res.MaxClaimLimit <= 0 {
// 		return 0, echo.NewHTTPError(400, "Voucher telah melebihi batas penggunaan")
// 	}
// 	// validasi tanggal
// 	now := time.Now()
// 	date := res.EndDate.Before(now)
// 	if !date {
// 		return 0, echo.NewHTTPError(400, "Telah melewati batas tanggal penggunaan voucher")
// 	}
// 	// validasi minimal belanjaan
// 	if res.MinimumPurchase > productCost {
// 		return 0, echo.NewHTTPError(400, "Total pembelian kurang untuk menggunakan voucher ini")
// 	}

// 	if res.VoucherType.Type == "Gratis Ongkir" {
// 		diskon = (shipCost * res.DiscountPercent) / 100
// 	}

// 	if res.VoucherType.Type == "Diskon Belanja" {
// 		diskon = (productCost * res.DiscountPercent) / 100

// 		if diskon > res.MaximumDiscount {
// 			diskon = res.MaximumDiscount
// 		}
// 	}

//		return diskon, nil
//	}
func (tu *transactionUsecase) ShippingOptions(ship *er.RajaongkirRequest) (interface{}, error) {

	res, err := rajaongkir.ShippingOptions(ship)
	if err != nil {
		return nil, err
	}

	return res, nil

}
