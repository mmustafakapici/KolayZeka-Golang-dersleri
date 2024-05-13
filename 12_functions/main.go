package main

import "fmt"

func main() {
	toplam1()
	plus(4, 5)
	result1 := plus2(8, 9, 6, 7)
	println("plus2fonk: ", result1)

	toplam, carpım := multi_islem(2, 1, 3, 4)

	fmt.Println("toplam : ", toplam, "carpım : ", carpım)

	sum(1, 2)

	sum(34, 34, 45, 56, 56)

	ints := []int{1, 2, 3, 4}

	sum(ints...)

	ret := sum_ret(ints...)

	println("ret : ", ret)
}

// paramete almayan
func toplam1() {
	var a = 2
	var b = 3
	println(a + b)
}

// değer döndürmeyen
func plus(a int, b int) {
	println("ikinci fonksion : ", a+b)
}

// değer döndüren fonk.
func plus2(a int, b int, c int, d int) int {
	return a + b + c + d
}

// multiple return values
func multi_islem(a int, b int, c int, d int) (int, int) {

	var x = a + b
	var y = c * d

	return x, y
}

// variadic fonksiyonlar

func sum(nums ...int) {

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)

}

func sum_ret(nums ...int) int {

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
