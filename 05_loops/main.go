package main

func main() {
	i := 0
	for i <= 4 {
		println(i)
		i = i + 1
	}

	println("ikinci:")

	for i := 1; i < 11; i = i + 2 {
		println(i)
	}

	for {
		println("loop")
		break
	}

	for n := range []string{"adana", "mersin", "2", "3", "6"} {
		println(n)
	}

	for n := range []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2} {
		if n%2 == 0 {
			println("index çifttir : atlandı ")
			continue
		}
		println("index tektir:", n)
	}

}
