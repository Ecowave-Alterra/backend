package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
func ShippingOptions(ship *et.ShippingRequest) (interface{}, error) {

	// malang kota
	alamatPengirim := "256"
	destination := ship.CityId
	weight := strconv.FormatUint(uint64(ship.Weight), 10)
	courier := []string{"jne", "pos", "tiki"}

	var result []et.ShippingResponse

	for _, val := range courier {
		url := "https://api.rajaongkir.com/starter/cost"

		log.Println(val)

		payloadStrings := fmt.Sprintf("origin=%s&destination=%s&weight=%s&courier=%s",
			alamatPengirim,
			destination,
			weight,
			val,
		)

		payload := strings.NewReader(payloadStrings)

		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("key", "8bb5248063ed493d90aac0311f8a3edb")
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseData et.ShippingResponse
		if err := json.Unmarshal(body, &responseData); err != nil {
			echo.NewHTTPError(500, "Can't Unmarshal JSON")
		}

		result = append(result, responseData)
	}

	return result, nil

}
