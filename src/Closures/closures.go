package main

import "fmt"

func multiBy2() func() int {
	i := 1
	return func() int {
		i *= 2
		return i
	}
}

func main() {

	nextMulti := multiBy2()

	fmt.Println(nextMulti()) // 2
	fmt.Println(nextMulti()) // 4
	fmt.Println(nextMulti()) // 8

	newNextMulti := multiBy2()
	fmt.Println(newNextMulti()) // 2
}
