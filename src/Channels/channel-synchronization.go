package main

import (
	"fmt"
	"time"
)

func worker(done chan string) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- "DONE"
}

func main() {

	done := make(chan string, 1)
	go worker(done)

	<-done //Block until worker is finished
}
