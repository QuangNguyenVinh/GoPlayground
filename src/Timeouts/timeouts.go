package main

import (
	"fmt"
	"time"
)

func worker(id int, input chan int, output chan int, timeout time.Duration) {
	fmt.Printf("Worker %d started\n", id)
	for {
		select {
		case job := <-input:
			fmt.Printf("Worker %d received job %d\n", id, job)
			time.Sleep(time.Duration(job) * time.Second)
			output <- job * 2
		case <-time.After(timeout):
			fmt.Printf("Worker %d timed out\n", id)
			return
		}
	}
}

func main() {
	input := make(chan int)
	output := make(chan int)
	timeout := time.Duration(3 * time.Second)

	go worker(1, input, output, timeout)
	go worker(2, input, output, 2*timeout)

	for i := 1; i <= 5; i++ {
		select {
		case outputValue := <-output:
			fmt.Printf("Result: %d\n", outputValue)
		case <-time.After(timeout):
			fmt.Println("Main timed out")
			return
		}
	}

	fmt.Println("Done")
}

