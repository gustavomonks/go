package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// here we can use i for index if we dont want to use the index we can change i for _.
	for i, num := range nums {
		sum += num
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println(sum)

	// range on map iterates over key/values pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
