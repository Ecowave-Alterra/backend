package binderbyte

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	etr "github.com/berrylradianh/ecowave-go/modules/entity/tracking"

	"github.com/labstack/echo/v4"
)

func Tracking(resi string, courier string) (interface{}, error) {
	var result etr.Tracking

	url := "https://api.binderbyte.com/v1/track?"
	param := fmt.Sprintf("api_key=%s&courier=%s&awb=%s",
		os.Getenv("BINDERBYTE_KEY"),
		courier,
		resi,
	)

	url = url + param

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		echo.NewHTTPError(500, "Can't Unmarshal JSON")
	}

	return result, nil
}
