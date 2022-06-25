package function

import (
	"fmt"
)

func helloWithFilter(name string, filter func(string) string) {

	fmt.Println("Hello", filter(name))

}

func spamFilter(Name string) string {

	if Name == "Anjing" {
		return "***"
	} else {
		return Name
	}

}

func callfilter() {

	helloWithFilter("Anjing", spamFilter)
}
