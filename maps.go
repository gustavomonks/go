package main

import "fmt"

func main() {
	//we use map when we want to access values by a key
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	a := map[string]int{"one": 1, "two": 2}
	fmt.Println(a)
}
