package main

func main() {
	nextInt := intSeq()
	println(nextInt())
	println(nextInt())
	println(nextInt())

	newInt := intSeq()
	println(newInt())
	println(nextInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		println("şu anda i değeri : ", i)
		i++
		return i
	}

}
