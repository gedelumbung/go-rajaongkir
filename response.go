package rajaongkir

type (
	/*ProvinceResponse : Response for province */
	ProvinceResponse struct {
		Rajaongkir struct {
			Status  RajaongkirStatus `json:"status"`
			Results Province         `json:"results"`
		} `json:"rajaongkir"`
	}

	/*ProvincesResponse : Response for provinces */
	ProvincesResponse struct {
		Rajaongkir struct {
			Status  RajaongkirStatus `json:"status"`
			Results []Province       `json:"results"`
		} `json:"rajaongkir"`
	}

	/*CityResponse : Response for city */
	CityResponse struct {
		Rajaongkir struct {
			Status  RajaongkirStatus `json:"status"`
			Results City             `json:"results"`
		} `json:"rajaongkir"`
	}

	/*CitiesResponse : Response for cities */
	CitiesResponse struct {
		Rajaongkir struct {
			Status  RajaongkirStatus `json:"status"`
			Results []City           `json:"results"`
		} `json:"rajaongkir"`
	}
)
