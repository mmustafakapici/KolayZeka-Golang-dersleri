package main

func main() {
	println(facto(2))
	println(facto(3))
	println(facto(4))
	println(facto(5))
	println(fib(1))

}

func facto(n int) int {
	if n == 0 {
		return 1
	}
	return n * facto(n-1)
}

// 1 2 3 4 5 6 7
// 1 1 2 3 5 8 13
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
