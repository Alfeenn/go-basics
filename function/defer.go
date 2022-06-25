package function

import (
	"fmt"
)

func Thx() {
	fmt.Println("Thx for ur submit")
}

func sayHello(value int) {
	defer Thx()
	fmt.Println("Your submit:")
	x := 10 / value
	fmt.Println("result", x)

}

func Main() {

	sayHello(0)

}
