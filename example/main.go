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
	cities := raja.Cities()
	fmt.Println(cities.Result)
	fmt.Println(cities.Error)
	cities = raja.CitiesByProvince(2)
	fmt.Println(cities.Result)
	fmt.Println(cities.Error)
	city := raja.City(23)
	fmt.Println(city.Result)
	fmt.Println(city.Error)
}
