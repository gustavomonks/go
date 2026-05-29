package main

import "fmt"

func loops() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for n := range 6 {
		// continue to the next iteration of the loop
		if n%2 == 0 {
			continue
		}
		fmt.Println(n + 1)
	}

}
