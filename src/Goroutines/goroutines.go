package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") //Sync

	go f("goroutine") //Async

	go func(msg string) {
		fmt.Println(msg)
	}("going") //Async

	go f("goroutine 2") //Async

	time.Sleep(time.Second) //Sleep until goroutines are finished. Otherwise main goroutine will exit and stop other goroutine.
	fmt.Println("done")
}
