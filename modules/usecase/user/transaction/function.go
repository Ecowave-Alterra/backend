package transaction

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/berrylradianh/ecowave-go/helper/hash"
	mdtrns "github.com/berrylradianh/ecowave-go/helper/midtrans"
	"github.com/berrylradianh/ecowave-go/helper/rajaongkir"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	em "github.com/berrylradianh/ecowave-go/modules/entity/midtrans"
	er "github.com/berrylradianh/ecowave-go/modules/entity/rajaongkir"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (string, string, error) {
	var productCost float64

	for _, cost := range transaction.TransactionDetails {
		stock, err := tu.transactionRepo.GetStock(cost.ProductId)
		if err != nil {
			return "", "", err
		}

		if stock < cost.Qty {
			return "", "", errors.New("Qty melebihi stock")
		}
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

	if err := vld.Validation(transaction); err != nil {
		return "", "", err
	}
	err = tu.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return "", "", err
	}
	return redirectUrl, transId, nil
}
func (tu *transactionUsecase) MidtransNotifications(midtransRequest *em.MidtransRequest) error {

	Key := hash.Hash(midtransRequest.OrderId, midtransRequest.StatusCode, midtransRequest.GrossAmount)

	if Key != midtransRequest.SignatureKey {
		log.Println(midtransRequest.SignatureKey)
		return errors.New("Invalid Transaction")
	}

	transaction := et.Transaction{
		TransactionId: midtransRequest.OrderId,
		PaymentStatus: midtransRequest.TransactionStatus,
		PaymentMethod: midtransRequest.PaymentType,
	}
	if midtransRequest.TransactionStatus == "settlement" {
		transaction.StatusTransaction = "Dikemas"
		transaction.PaymentStatus = midtransRequest.TransactionStatus
	}
	if midtransRequest.TransactionStatus == "pending" {
		transaction.PaymentStatus = midtransRequest.TransactionStatus
	}
	if midtransRequest.TransactionStatus == "expire" {
		transaction.StatusTransaction = "Dibatalkan"
		transaction.CanceledReason = "pembayaran kadaluarsa"
		transaction.PaymentStatus = midtransRequest.TransactionStatus
	}
	if midtransRequest.TransactionStatus == "failure" {
		transaction.StatusTransaction = "Dibatalkan"
		transaction.CanceledReason = "pembayaran gagal"
		transaction.PaymentStatus = midtransRequest.TransactionStatus
	}

	err := tu.transactionRepo.UpdateTransaction(transaction)
	if err != nil {
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

func (tu *transactionUsecase) ShippingOptions(ship *er.RajaongkirRequest) (interface{}, error) {

	if err := vld.Validation(ship); err != nil {
		return nil, err
	}

	res, err := rajaongkir.ShippingOptions(ship)
	if err != nil {
		return nil, err
	}

	return res, nil

}
