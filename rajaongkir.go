package rajaongkir

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	ErrNotFound   = errors.New("Data not found")
	ErrInvalidKey = errors.New("Invalid Key")
)

type Result struct {
	Result interface{}
	Error  error
}

type RajaongkirStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type RajaOngkir struct {
	Key         string
	Url         string
	AccountType string
}

type Province struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

type City struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type ProvinceResponse struct {
	Rajaongkir struct {
		Status  RajaongkirStatus `json:"status"`
		Results Province         `json:"results"`
	} `json:"rajaongkir"`
}

type ProvincesResponse struct {
	Rajaongkir struct {
		Status  RajaongkirStatus `json:"status"`
		Results []Province       `json:"results"`
	} `json:"rajaongkir"`
}

type CityResponse struct {
	Rajaongkir struct {
		Status  RajaongkirStatus `json:"status"`
		Results City             `json:"results"`
	} `json:"rajaongkir"`
}

type CitiesResponse struct {
	Rajaongkir struct {
		Status  RajaongkirStatus `json:"status"`
		Results []City           `json:"results"`
	} `json:"rajaongkir"`
}

func Init(key, url, accountType string) *RajaOngkir {
	return &RajaOngkir{key, url, accountType}
}

func (r *RajaOngkir) Province(id int) Result {
	model := &ProvinceResponse{}
	provinceID := strconv.Itoa(id)
	err := r.call(http.MethodGet, r.Url+`/`+r.AccountType+`/province?id=`+provinceID, model)
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

func (r *RajaOngkir) Provinces() Result {
	model := &ProvincesResponse{}
	err := r.call(http.MethodGet, r.Url+`/`+r.AccountType+`/province`, model)
	statusCode := model.Rajaongkir.Status.Code
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

func (r *RajaOngkir) Cities() Result {
	model := &CitiesResponse{}
	err := r.call(http.MethodGet, r.Url+`/`+r.AccountType+`/city`, model)
	statusCode := model.Rajaongkir.Status.Code
	if statusCode == 400 {
		return Result{Error: ErrInvalidKey}
	}
	if err != nil {
		return Result{Error: err}
	}
	return Result{Result: model}
}

func (r *RajaOngkir) CitiesByProvince(provinceId int) Result {
	model := &CitiesResponse{}
	provinceID := strconv.Itoa(provinceId)
	err := r.call(http.MethodGet, r.Url+`/`+r.AccountType+`/city?province=`+provinceID, model)
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

func (r *RajaOngkir) City(id int) Result {
	model := &CityResponse{}
	cityID := strconv.Itoa(id)
	err := r.call(http.MethodGet, r.Url+`/`+r.AccountType+`/city?id=`+cityID, model)
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
