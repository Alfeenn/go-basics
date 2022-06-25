package packages

import (
	"fmt"

	"time"
)

func Timme() {

	layout := "Jan 2,2006 at 3:04PM" //Harus set up base layout sesuai dengan aturan!
	Time, error := time.Parse(layout, "May 12,2020 at 1:00AM")
	if error == nil {

		fmt.Println("Waktu saat ini :", Time)
	} else {
		fmt.Println(error.Error())
	}
	loc, _ := time.LoadLocation("Europe/Berlin")
	const timelayout = "2006-Jan-02 at 3:04PM"
	time, _ := time.ParseInLocation(timelayout, "2020-May-12 at 12:22PM", loc)

	fmt.Println(time)


	
}
