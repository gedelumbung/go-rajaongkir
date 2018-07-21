package rajaongkir

type (
	/*RajaOngkir : Rajaongkir Struct */
	RajaOngkir struct {
		Key         string
		URL         string
		AccountType string
	}

	/*Result : Result Struct */
	Result struct {
		Result interface{}
		Error  error
	}

	/*RajaongkirStatus : RajaongkirStatus Struct */
	RajaongkirStatus struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}

	/*Province : Province Struct */
	Province struct {
		ProvinceID string `json:"province_id"`
		Province   string `json:"province"`
	}

	/*City : City Struct */
	City struct {
		CityID     string `json:"city_id"`
		ProvinceID string `json:"province_id"`
		Province   string `json:"province"`
		Type       string `json:"type"`
		CityName   string `json:"city_name"`
		PostalCode string `json:"postal_code"`
	}
)
