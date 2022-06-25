package packages

import (
	"flag"
	"fmt"
)

func Flag() {

	host := flag.String("Host", "localhost", "Put value")
	username := flag.String("Username", "root", "Put value")
	password := flag.String("Password", "root", "Put value")
	number := flag.Int("Number", 123, "Put value")

	flag.Parse()
	fmt.Println("Host :", *host)
	fmt.Println("User :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Number :", *number)

}
