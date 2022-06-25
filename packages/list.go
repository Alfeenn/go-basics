package packages

import (
	"container/list"
	"fmt"
)

func main() {
	var NewE *list.Element
	data := list.New() //membuat data linked list baru

	data.PushBack("Alfian") //memasukkan data ke list paling belakang
	data.PushBack(2)
	data.PushBack("George")
	data.PushFront("Mike") //memasukkan data  ke list paling depan

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)

	}

	for NewE = data.Front(); NewE.Value != "Mike"; NewE = NewE.Next() {
	}
	data.MoveToBack(NewE) //pengecekan value pada fungsi for,jika terdapat value yang dituju maka move value ke belakang

	for NewE = data.Front(); NewE != nil; NewE = NewE.Next() {
		fmt.Println(NewE.Value)
	}

}
