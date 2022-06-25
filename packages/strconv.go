package packages

import (
	"flag"
	"fmt"
	"strconv"
)

func Strconv() {

	// plaintext := "1234"

	password := flag.String("password", "2110", "Put password here")
	flag.Parse()

	number, err := strconv.ParseInt(*password, 0, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

}
