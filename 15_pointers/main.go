package main

func main() {

	var sayı1 float64 = 32
	var sayı2 float64 = 32
	i := 4
	int2 := 6
	println("initial: ", i)
	zeroval(i)
	println("zeroval: ", i)
	zeroptr(&i)
	println("zeroptr: ", i)
	println("pointer: ", &i)

	println("pointer: ", &int2)

	println("sayı1 adresi: ", &sayı1)
	println("sayı2 adresi: ", &sayı2)

	var str1 string = "123456789"
	var str2 string = "abcd"

	println("str1 adresi: ", &str1)
	println("str2 adresi: ", &str2)

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
