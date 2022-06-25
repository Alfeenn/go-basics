package function

import (
	"fmt"
)

func Greeting(name string) string {

	return "Hello" + name

}

//function as value -> can turn function to value  ex : print(variable(function as value))
func Greet() {

	goGreet := Greeting
	fmt.Println(goGreet("mike"))

}
