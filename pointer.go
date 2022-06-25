package main

import (
	"fmt"
)

type Address struct {
	city, province, country string
}

func address1() Address {

	currentAddress := Address{city: "serang"}
	currentAddress.province = "Banten"
	currentAddress.country = "Indonesia"
	return currentAddress
}

func City() {

	city1 := address1()
	city2 := &city1
	city3 := &city1
	fmt.Println(city1)

	*city2 = Address{"Cilegon", "Banten", "Indonesia"}
	fmt.Println(city1)
	fmt.Println(city2)
	fmt.Println(city3)

}
