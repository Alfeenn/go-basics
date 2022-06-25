package main

import (
	"fmt"
	"os"
)

func OS() {

	fmt.Println(os.Args)
	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)
	}
}
