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

	var s = make([]string, 3)
	fmt.Println(s)
	// append only works when we use slice, because arrays have a fixed sized
	s = append(s, "a")
	fmt.Println(s)
	fmt.Println(len(s))
	s[1] = "b"
	s[0] = "c"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(s[3:4]) // create a slice starting 3 and end before 4

}
