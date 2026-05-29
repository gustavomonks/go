package main

import "fmt"

func main() {

	x := 2
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//switch to compares types instead of values.
	whatAmI := func(i interface{}) { //inteface mean we accpet any type
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println(t)
		}
	}

	whatAmI(true)
	whatAmI("oi")
	whatAmI(1)
}
