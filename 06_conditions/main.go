package main

func main() {
	n := 2
	if n%2 == 0 {
		println("n çift sayıdır n : ", n)
	} else {
		println("n tek sayıdır n : ", n)
	}

	println("yeni durum")

	if 8%4 == 0 {
		println("sekiz dörde tam bölünür")
	}

	a := 4
	b := 4
	if a%2 == 0 || b%2 == 0 {
		println("a ya da b' den az biri çifttir")
	}

	num := 0
	if num > 0 {
		println(num, " sıfırdan büyüktür")
	} else if num < 0 {
		println(num, "sıfırdan küçüktür")
	} else {
		println(num, "sıfıra eşittir")
	}
}
