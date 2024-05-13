package main

import "fmt"

func main() {
	r := rect{5, 10}
	fmt.Println("area : ", r.area())
	fmt.Println("perimeter : ", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("cevre:", rp.perimeter())
	r2 := rect{50, 10}
	fmt.Println("toplam", r2.degerler_top())
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func (r rect) degerler_top() int {
	return r.height + r.width
}
