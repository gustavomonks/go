package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r rectangle) Area() int {
	return r.height * r.width
}

func main() {
	rect := rectangle{5, 2}
	fmt.Println(rect.Area())

}
