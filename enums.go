package main

import "fmt"

// iota automatically generates incrementing integer, values startind from 0 within constant declarations
// iota is kind of an index iterator

type StatusResponse int

const (
	PENDING  StatusResponse = iota //0
	ACCEPTED                       // 1
	REFUSED                        // 2

)

var statusName = map[StatusResponse]string{
	PENDING:  "Pending",
	ACCEPTED: "Accepted",
	REFUSED:  "Refused",
}

// this method is used to translate every StatusResponse to String()
func (sr StatusResponse) String() string {
	return statusName[sr]
}

func main() {
	fmt.Println(ACCEPTED)
}
