package main

import (
	"fmt"
	"time"
)

// Timer fires only once, while a ticker fires repeatedly.
func main() {
	fmt.Println("Start")

	// Create a timer that fires after 2 seconds
	timer := time.After(2 * time.Second)
	<-timer
	fmt.Println("Timer expired")

	// Create a ticker that ticks every second for 5 seconds
	ticker := time.Tick(1 * time.Second)
	for i := 0; i < 5; i++ {
		<-ticker
		fmt.Println("Tick")
	}
}
