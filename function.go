package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

// sum accepts a variable number of integers
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	fmt.Println(vals())
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)
	sum(1, 2, 3, 4)

}
