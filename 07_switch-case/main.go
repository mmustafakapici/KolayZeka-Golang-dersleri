package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	switch i {
	case 1:
		println("sayımız 1")
	case 2:
		println("sayımız 2")
	case 3:
		println("sayımız 3")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("weekend")

	default:
		println("weekday")

	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		println("öğleden önce")

	default:
		println("öğleden sonra")

	}

	whatAmI := func(i interface{}) {
		switch degisken_tipi := i.(type) {
		case bool:
			println("variable is bool")
		case int:
			println("variable is int")
		default:
			fmt.Printf("I dont know maybe %T\n", degisken_tipi)
		}
	}
	whatAmI(true)
	whatAmI(345)
	whatAmI("kolay zeka")
	whatAmI(34234.04324)

}
