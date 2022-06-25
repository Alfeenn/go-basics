package main

import (
	"fmt"
	"reflect"
)

type People struct {
	a1 string `alias : "num1"`
	a2 string `alias : ""`
	a3 string 
}

func main() {

	p1 := People{}
	sample := reflect.TypeOf(p1)

	for i:= 0; i<sample.NumField();i++{


		}

}