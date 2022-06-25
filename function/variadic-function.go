package function

import (
	"fmt"
)

func sumAll(numbers ...int) int { // ( ...int)-> variadic number = bisa memasukkan lebih dari 1 parameter tanpa membuat array

	total := 0
	for _, sum := range numbers {
		total += sum
	}
	return total
}

func sum() {

	total := sumAll()
	fmt.Println(total)
}
