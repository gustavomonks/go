package main

import "fmt"

type person struct {
	name string
	age  int
}

// &p returns a pointer to p.
// The function returns *person, so we return the address of the struct.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(newPerson("Gustavo"))
	fmt.Println(person{"Joao", 25})
	fmt.Println(&person{"Carlos", 17})

	s := person{"Giovanna", 25}
	fmt.Println(s.name)
	fmt.Println(s.age)
	s.age = 27
	fmt.Println(s)

	dog := struct {
		name   string
		isGood bool
	}{
		"Vanessa",
		true,
	}

	fmt.Println(dog)
}
