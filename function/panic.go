package function

import (
	"fmt"
)

func End() {
	fmt.Println("Application exit")
}

func runApp(error bool) {
	defer End()
	if error {
		panic("Something wrong")
	}
	fmt.Println("app running. . .")

}

func run() {

	runApp(true)

}
