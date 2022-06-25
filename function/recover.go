package function

import (
	"fmt"
)

func Endapp() {
	msg := recover()
	if msg != nil {
		fmt.Println("Error msg:", msg)
	}
	fmt.Println("Application exit")
}

func Runapp(error bool) {
	defer Endapp()
	if error {
		panic("Apps error")
	}
	fmt.Println("Apps running")

}

func Test() {

	Runapp(false)

}
