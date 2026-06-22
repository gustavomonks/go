package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func AddInt(a int, b int) int {
	return a + b
}

func AddFloat(a float64, b float64) float64 {
	return a + b
}

func Add[T int | float64](a T, b T) T {
	return a + b
}

// -----

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type UserGenerics[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func main() {

	fmt.Println(AddFloat(2, 4.5))
	fmt.Println(AddInt(2, 4))
	//fmt.Println(AddInt(2.23, 1)) //wont work
	fmt.Println(Add(2.523, 1))
	// here we pass what type we want for Data
	u := UserGenerics[int]{
		ID:   2,
		Name: "Gustavo",
		Data: 19,
	}
	fmt.Println(u)

}
