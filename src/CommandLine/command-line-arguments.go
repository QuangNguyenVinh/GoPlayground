package main

import (
	"fmt"
	"os"
)

// go build command-line-arguments.go
// command-line-arguments a b c d
func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
