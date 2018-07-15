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
