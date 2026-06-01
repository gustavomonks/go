package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(len(a))
	a[len(a)-1] = 2
	fmt.Println(a)

	b := [...]int{1, 2, 3}
	fmt.Println(b)

	var c = [2][3]int{
		{1, 2, 3},
		{3, 2, 1},
	}

	fmt.Println(c)
}
