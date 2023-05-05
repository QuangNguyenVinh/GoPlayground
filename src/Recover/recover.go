package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

//A recover can stop a panic from aborting the program and let it continue with execution instead.
func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
