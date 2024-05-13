package main

import "fmt"

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"Ann", 40})
	fmt.Println(newPerson("Jhon"))

	s := person{"Sean", 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{"Rex", true}

	fmt.Println(dog)
}

// structs
type person struct {
	name string // field
	age  int    // field
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
