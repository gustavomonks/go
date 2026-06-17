package main

import "fmt"

func elseif() {
	//there is n o ternaty if in Go

	const a int = 10
	if a > 15 {
		fmt.Println("greater than 15")
	} else if a < 15 {
		fmt.Println("less than 15")
	} else {
		fmt.Println("equal to 15")
	}

	var b int
	fmt.Println("Enter a number: ")
	fmt.Scan(&b) //the typed value will be store inside b

	if b > 15 {
		fmt.Println("greater than 15")
	} else if b < 15 {
		fmt.Println("less than 15")
	} else {
		fmt.Println("equal to 15")
	}
}
