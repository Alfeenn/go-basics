package function

import (
	"fmt"
)

func Callperson() {

	type person interface{}

	var Person = map[string]person{
		"\n1. Name\t": "Alfian",
		"\n2. Age\t":  "22",
		"\n3. Food\t": []string{"Meatballs", "Chicken", "Apple"},
	}

	fmt.Println(Person)

}
