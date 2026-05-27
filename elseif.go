package main

import "fmt"

func main() {
	//there is n o ternaty if in Go

	const a int = 10
	if a > 15 {
		fmt.Println("greater than 15")
	} else if a < 15 {
		fmt.Println("less than 15")
	} else {
		fmt.Println("equal to 15")
	}

}
