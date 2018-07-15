package main

import (
	"fmt"

	rajaongkir "github.com/gedelumbung/go-rajaongkir"
)

const (
	key         = "f5b38c7a0e7272989375aaf8b8eb2586"
	url         = "https://api.rajaongkir.com"
	accountType = "starter"
)

func main() {
	raja := rajaongkir.Init(key, url, accountType)
	provinces := raja.Provinces()
	fmt.Println(provinces.Result)
	fmt.Println(provinces.Error)
	province := raja.Province(10)
	fmt.Println(province.Result)
	fmt.Println(province.Error)
}
