package main

import "fmt"

//https://www.youtube.com/watch?v=W3xS52aR0dM
// solution for a coding challenge from Gary Clarke's

type Logger interface {
	Log() string
}

type ConsoleLogger struct {
	message string
}

type PrefixedLogger struct {
	message string
}

func (log ConsoleLogger) Log() string {
	return log.message
}

func (log PrefixedLogger) Log() string {
	return "[INFO] " + log.message
}

func runLog(log Logger) {
	fmt.Println(log.Log())
}

func main() {
	c := ConsoleLogger{"blabla"}
	p := PrefixedLogger{"some infos"}

	runLog(c)
	runLog(p)
}
