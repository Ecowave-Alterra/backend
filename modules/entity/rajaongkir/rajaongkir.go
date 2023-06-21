package rajaongkir

type RajaongkirRequest struct {
	Weight float32 `json:"Weight" validate:"required"`
	CityId string  `json:"CityId" validate:"required"`
}

type RajaongkirResponse struct {
	Rajaongkir struct {
		Results []struct {
			Code  string `json:"Code"`
			Name  string `json:"Name"`
			Costs []struct {
				Service     string `json:"Service"`
				Description string `json:"Description"`
				Cost        []struct {
					Value uint   `json:"Value"`
					Etd   string `json:"Etd"`
				} `json:"Cost"`
			} `json:"Costs"`
		} `json:"Results"`
	} `json:"Rajaongkir"`
}

type Results struct {
	Code        string
	Name        string
	Service     string
	Description string
	Value       uint
	Etd         string
}
