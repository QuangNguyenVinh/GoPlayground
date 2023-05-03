package main

import (
	"fmt"
	"time"
)

func main() {
	first_channel := make(chan string)
	second_channel := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		first_channel <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		second_channel <- "World"
	}()

	for i := 0; i < 2; i += 1 {
		//Wait on multiple channel operations.
		select {
		case first_msg := <-first_channel:
			fmt.Println(first_msg)
		case second_msg := <-second_channel:
			fmt.Println(second_msg)
		}
	}
}
