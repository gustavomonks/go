package main

import "fmt"

type Address struct {
	City  string
	State string
}

func (a Address) Full() string {
	return a.City + " " + a.State
}

type User struct {
	Name    string
	Address // we dont need to define type here
}

func main() {
	user := User{
		Name: "Gustavo",
		Address: Address{
			City:  "Araraquara",
			State: "SP",
		},
	}

	fmt.Println(user.Name)
	fmt.Println(user.City)
	fmt.Println(user.Full()) // we can call Address methods directly
}
