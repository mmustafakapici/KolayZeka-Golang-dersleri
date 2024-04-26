package main

import (
	"fmt"
	"math"
)

const s string = "sabit"
const n = 500000000
const d = 3e20 / n

func main() {

	//s = "degistirildi"
	println(s)
	println(n)
	println(d)
	println(int64(d))

	fmt.Println(math.Sin(n))

}
