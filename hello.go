package main

import "fmt"

func hello() {
	x := 10
	fmt.Println("Hello, World!")
	fmt.Println(x)
	x = 20
	fmt.Println(x)
	var a string = "2"
	fmt.Println(a)
	const b bool = true
	fmt.Println(!b)
}
