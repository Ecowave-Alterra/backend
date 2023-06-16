package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/berrylradianh/ecowave-go/helper/hash"
	mdtrns "github.com/berrylradianh/ecowave-go/helper/midtrans"
	em "github.com/berrylradianh/ecowave-go/modules/entity/midtrans"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

func (tu *transactionUsecase) CreateTransaction(transaction *et.Transaction) (string, error) {
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
		return "", err
	}

	err = tu.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return "", err
	}
	return redirectUrl, nil
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
func ShippingOptions(ship *et.ShippingRequest) (interface{}, error) {

	// malang kota
	alamatPengirim := "256"
	destination := ship.CityId
	weight := strconv.FormatUint(uint64(ship.Weight), 10)
	courier := []string{"jne", "pos", "tiki"}

	var result []et.ShippingResponse

	for _, val := range courier {
		url := "https://api.rajaongkir.com/starter/cost"

		payloadStrings := fmt.Sprintf("origin=%s&destination=%s&weight=%s&courier=%s",
			alamatPengirim,
			destination,
			weight,
			val,
		)

		payload := strings.NewReader(payloadStrings)

		key := os.Getenv("RAJAONGKIR_KEY")

		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("key", key)
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

// type Client snap.Client

// // New : this function will always be called when the Snap is initiated
// func (c *Client) New(serverKey string, env midtrans.EnvironmentType) {
// 	c.Env = env
// 	c.ServerKey = serverKey
// 	c.Options = &midtrans.ConfigOptions{}
// 	c.HttpClient = midtrans.GetHttpClient(env)
// }

// func getDefaultClient() Client {
// 	return Client{
// 		ServerKey:  midtrans.ServerKey,
// 		Env:        midtrans.Environment,
// 		HttpClient: midtrans.GetHttpClient(midtrans.Environment),
// 		Options: &midtrans.ConfigOptions{
// 			PaymentOverrideNotification: midtrans.PaymentOverrideNotification,
// 			PaymentAppendNotification:   midtrans.PaymentAppendNotification,
// 		},
// 	}
// }

// func (c Client) CreateTransactionWithMap(req *snap.RequestParamWithMap) (snap.ResponseWithMap, *midtrans.Error) {
// 	resp := snap.ResponseWithMap{}
// 	jsonReq, _ := json.Marshal(req)
// 	err := c.HttpClient.Call(
// 		http.MethodPost,
// 		fmt.Sprintf("%s/snap/v1/transactions", c.Env.SnapURL()),
// 		&c.ServerKey,
// 		c.Options,
// 		bytes.NewBuffer(jsonReq),
// 		&resp,
// 	)

// 	if err != nil {
// 		return resp, err
// 	}
// 	return resp, nil
// }
// func CreateTransactionWithMap(req *snap.RequestParamWithMap) (snap.ResponseWithMap, *midtrans.Error) {
// 	return getDefaultClient().CreateTransactionWithMap(req)
// }
// func (c Client) CreateTransactionTokenWithMap(req *snap.RequestParamWithMap) (string, *midtrans.Error) {
// 	var snapToken string
// 	resp, err := c.CreateTransactionWithMap(req)

// 	if err != nil {
// 		return snapToken, err
// 	}

// 	if token, found := resp["token"]; !found {
// 		return snapToken, &midtrans.Error{
// 			Message:    "Token field notfound",
// 			StatusCode: 0,
// 		}
// 	} else {
// 		snapToken = token.(string)
// 		return snapToken, nil
// 	}
// }

// // CreateTransactionTokenWithMap : Do `/transactions` API request to SNAP API to get Snap token with map as
// // body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func CreateTransactionTokenWithMap(req *snap.RequestParamWithMap) (string, *midtrans.Error) {
// 	return getDefaultClient().CreateTransactionTokenWithMap(req)
// }

// // CreateTransactionUrlWithMap : Do `/transactions` API request to SNAP API to get Snap redirect url with map as
// // body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func (c Client) CreateTransactionUrlWithMap(req *snap.RequestParamWithMap) (string, *midtrans.Error) {
// 	var redirectUrl string
// 	resp, err := c.CreateTransactionWithMap(req)

// 	if err != nil {
// 		return redirectUrl, err
// 	}

// 	if url, found := resp["redirect_url"]; !found {
// 		return redirectUrl, &midtrans.Error{
// 			Message:    "Error redirect_url field notfound in json response",
// 			StatusCode: 0,
// 		}
// 	} else {
// 		redirectUrl = url.(string)
// 		return redirectUrl, nil
// 	}
// }

// // CreateTransactionUrlWithMap : Do `/transactions` API request to SNAP API to get Snap redirect url with map
// // as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func CreateTransactionUrlWithMap(req *snap.RequestParamWithMap) (string, *midtrans.Error) {
// 	return getDefaultClient().CreateTransactionUrlWithMap(req)
// }

// // CreateTransaction : Do `/transactions` API request to SNAP API to get Snap token and redirect url with `snap.Request`
// // as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func (c Client) CreateTransaction(req *snap.Request) (*snap.Response, *midtrans.Error) {
// 	resp := &snap.Response{}
// 	jsonReq, _ := json.Marshal(req)
// 	err := c.HttpClient.Call(
// 		http.MethodPost,
// 		fmt.Sprintf("%s/snap/v1/transactions", c.Env.SnapURL()),
// 		&c.ServerKey,
// 		c.Options,
// 		bytes.NewBuffer(jsonReq),
// 		resp,
// 	)

// 	if err != nil {
// 		return resp, err
// 	}
// 	return resp, nil
// }

// // CreateTransaction : Do `/transactions` API request to SNAP API to get Snap token and redirect url with `snap.Request`
// // as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func CreateTransaction(req *snap.Request) (*snap.Response, *midtrans.Error) {
// 	return getDefaultClient().CreateTransaction(req)
// }

// // CreateTransactionToken : Do `/transactions` API request to SNAP API to get Snap token with `snap.Request` as
// // body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func (c Client) CreateTransactionToken(req *snap.Request) (string, *midtrans.Error) {
// 	var snapToken string
// 	resp, err := c.CreateTransaction(req)
// 	if err != nil {
// 		return snapToken, err
// 	}

// 	if resp.Token != "" {
// 		snapToken = resp.Token
// 	}
// 	return snapToken, nil
// }

// // CreateTransactionToken : Do `/transactions` API request to SNAP API to get Snap token with `snap.Request` as
// // body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func CreateTransactionToken(req *snap.Request) (string, *midtrans.Error) {
// 	return getDefaultClient().CreateTransactionToken(req)
// }

// // CreateTransactionUrl : Do `/transactions` API request to SNAP API to get Snap redirect url with `snap.Request`
// // as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func (c Client) CreateTransactionUrl(req *snap.Request) (string, *midtrans.Error) {
// 	var redirectUrl string
// 	resp, err := c.CreateTransaction(req)
// 	if err != nil {
// 		return redirectUrl, err
// 	}

// 	if resp.RedirectURL != "" {
// 		redirectUrl = resp.RedirectURL
// 	}
// 	return redirectUrl, nil
// }

// // CreateTransactionUrl : Do `/transactions` API request to SNAP API to get Snap redirect url with `snap.Request`
// // as body parameter, will be converted to JSON, more detail refer to: https://snap-docs.midtrans.com
// func CreateTransactionUrl(req *snap.Request) (string, *midtrans.Error) {
// 	return getDefaultClient().CreateTransactionUrl(req)
// }
