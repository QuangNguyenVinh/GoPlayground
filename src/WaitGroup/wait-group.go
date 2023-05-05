package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the WaitGroup counter
		go func(id int) {
			defer wg.Done() // decrement the WaitGroup counter when the goroutine exits
			//Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
			//This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
			fmt.Printf("Worker %d started\n", id)
			time.Sleep(2 * time.Second) // simulate some work
			fmt.Printf("Worker %d finished\n", id)
		}(i)
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All workers finished")
}
