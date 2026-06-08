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

// anonymous funciton
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	fmt.Println(vals())
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)
	sum(1, 2, 3, 4)
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(intSeq()()) //we use two () because the first call returns a func and the second call execute that func

}
