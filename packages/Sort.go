package packages

import (
	"fmt"
	"sort" //package sort untuk mensortit data
)

type User struct { //membuat struct data untuk disortir
	Name string
	Age  int
}

type UserSlice []User // membuat alias struct

func (value UserSlice) Len() int { //hitung panjang slice

	return len(value)
}

func (user UserSlice) Less(i, j int) bool { //sortir data
	return user[i].Age < user[j].Age
}

func (user UserSlice) Swap(i, j int) { //swap data
	user[i], user[j] = user[j], user[i]
}

func Sort() {

	user := []User{ //membuat variable slice struct user
		{"Alfian", 20},
		{"ian", 1},
		{"A", 2},
	}

	sort.Sort(UserSlice(user)) //func sort user (*Note : argument harus implement package sort )
	fmt.Println(user)

}
