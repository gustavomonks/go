package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	a := map[string]int{"one": 1, "two": 2}
	fmt.Println(a)
}
