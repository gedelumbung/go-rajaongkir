package rajaongkir

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	/*ErrNotFound : Default error when data not found */
	ErrNotFound = errors.New("Data not found")
	/*ErrInvalidKey : Default error when api key is invalid */
	ErrInvalidKey = errors.New("Invalid Key")
)

/*Init : Initial for rajaongkir library*/
/**
 * @params (key string, url string, accountType string)
 * @return *Rajaongkir
 */
func Init(key, url, accountType string) *RajaOngkir {
	return &RajaOngkir{key, url, accountType}
}

/*Province : Get selected/detail province by province id */
/**
 * @params (id int)
 * @return Result
 */
func (r *RajaOngkir) Province(id int) Result {
	model := &ProvinceResponse{}
	provinceID := strconv.Itoa(id)
	err := r.call(http.MethodGet, r.URL+`/`+r.AccountType+`/province?id=`+provinceID, model)
	statusCode := model.Rajaongkir.Status.Code
	provinceID = model.Rajaongkir.Results.ProvinceID
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if provinceID == "" {
		return Result{Error: ErrNotFound}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

/*Provinces : Get all provinces data */
/**
 * @return Result
 */
func (r *RajaOngkir) Provinces() Result {
	model := &ProvincesResponse{}
	err := r.call(http.MethodGet, r.URL+`/`+r.AccountType+`/province`, model)
	statusCode := model.Rajaongkir.Status.Code
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

/*Cities : Get all cities data */
/**
 * @return Result
 */
func (r *RajaOngkir) Cities() Result {
	model := &CitiesResponse{}
	err := r.call(http.MethodGet, r.URL+`/`+r.AccountType+`/city`, model)
	statusCode := model.Rajaongkir.Status.Code
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

/*CitiesByProvince : Get all cities by province id */
/**
 * @params (provinceID int)
 * @return Result
 */
func (r *RajaOngkir) CitiesByProvince(provinceID int) Result {
	model := &CitiesResponse{}
	id := strconv.Itoa(provinceID)
	err := r.call(http.MethodGet, r.URL+`/`+r.AccountType+`/city?province=`+id, model)
	statusCode := model.Rajaongkir.Status.Code
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if len(model.Rajaongkir.Results) == 0 {
		return Result{Error: ErrNotFound}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

/*City : Get city by city id */
/**
 * @params (id int)
 * @return Result
 */
func (r *RajaOngkir) City(id int) Result {
	model := &CityResponse{}
	cityID := strconv.Itoa(id)
	err := r.call(http.MethodGet, r.URL+`/`+r.AccountType+`/city?id=`+cityID, model)
	statusCode := model.Rajaongkir.Status.Code
	cityID = model.Rajaongkir.Results.CityID
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if cityID == "" {
		return Result{Error: ErrNotFound}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

func (r *RajaOngkir) call(method, endpoint string, model interface{}) error {
	req, err := http.NewRequest(method, endpoint, nil)
	req.Header.Add("key", r.Key)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &model)
	return err
}
