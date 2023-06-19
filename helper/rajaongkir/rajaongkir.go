package rajaongkir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	er "github.com/berrylradianh/ecowave-go/modules/entity/rajaongkir"
	"github.com/labstack/echo/v4"
)

func ShippingOptions(ship *er.RajaongkirRequest) (interface{}, error) {
	// malang kota
	alamatPengirim := "256"
	destination := ship.CityId
	weight := strconv.FormatUint(uint64(ship.Weight), 10)
	courier := []string{"jne", "pos", "tiki"}

	var rajaongkirReq []er.RajaongkirResponse
	var result []er.Results

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

		var responseData er.RajaongkirResponse
		if err := json.Unmarshal(body, &responseData); err != nil {
			echo.NewHTTPError(500, "Can't Unmarshal JSON")
		}

		rajaongkirReq = append(rajaongkirReq, responseData)
	}

	for _, val := range rajaongkirReq {
		for _, v := range val.Rajaongkir.Results {
			res := er.Results{
				Code: v.Code,
				Name: v.Name,
			}
			for _, c := range v.Costs {
				res.Service = c.Service
				res.Description = c.Description
				for _, cs := range c.Cost {
					res.Value = cs.Value
					res.Etd = cs.Etd
				}
			}

			result = append(result, res)
		}
	}
	return result, nil
}
