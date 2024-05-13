package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	s = make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	//s[3] = "d"
	s = append(s, "d", "e", "f")
	fmt.Println("slice: ", s, "len: ", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c slice: ", c)

	l1 := s[2:5]
	l2 := s[:5]
	l3 := s[2:]

	fmt.Println("l1:", l1, "l2: ", l2, "l3: ", l3)

	strings := []string{"g", "h", "j"}

	fmt.Println("declared: ", strings)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i - j
		}
	}
	fmt.Println("iki boyut:", twoD)
}
