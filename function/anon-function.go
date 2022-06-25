package function

import (
	"fmt"
)

// type Blacklist func(name string) bool

func registerUser(name string, blacklist func(name string) bool) {
	if blacklist(name) {
		fmt.Println("Username blocked:", name)
	} else {
		fmt.Println("Welcome", name)
	}

}

func call() {
	User := func(name string) bool {
		return name == "A" //->Anonymous function
	}

	registerUser("Admin", User)
	registerUser("Anjing", User)

}
