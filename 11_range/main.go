package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum : ", sum)

	println("range ile döngü")
	for i, num := range nums {
		fmt.Println("index: ", i, "num : ", num)
	}

	println("for i  ile döngü")

	for i := 0; i < len(nums); i++ {
		fmt.Println("index: ", i, "num : ", nums[i])
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}
	println("ascii")

	for i, c := range "<> +$%&()[]{} ;:,. " {
		fmt.Println(i, c)
	}

}
