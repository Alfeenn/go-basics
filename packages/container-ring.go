package packages

import (
	"container/ring"
	"fmt"
)

func Ring() {
	var data *ring.Ring = ring.New(5) //tipe data ring ,apabila sudah sampai max data makan akan kembali ke data awal

	for i := 0; i < data.Len(); i++ {

		data.Value = i
		data = data.Next()
	}

	data.Do(func(number interface{}) {

		fmt.Println(number)
	})

}
