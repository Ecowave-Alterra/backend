package tracking

type Tracking struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Summary struct {
			Awb     string `json:"awb"`
			Courier string `json:"courier"`
			Service string `json:"service"`
			Status  string `json:"status"`
			Date    string `json:"date"`
			Desc    string `json:"desc"`
			Amount  string `json:"amount"`
			Weight  string `json:"weight"`
		} `json:"summary"`
		Detail struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Shipper     string `json:"shipper"`
			Receiver    string `json:"receiver"`
		} `json:"detail"`
		History []struct {
			Date     string `json:"date"`
			Desc     string `json:"desc"`
			Location string `json:"location"`
		} `json:"history"`
	} `json:"data"`
}
