package main

import "fmt"

// Main goroutine
func main() {

	//Channels are the pipes that connect concurrent goroutines.
	//You can send values into channels from one goroutine and receive those values into another goroutine.

	messages := make(chan string) //Syntax create channel: make(chan val-type).

	go func() { messages <- "ping" }() //Anonymous function and goroutine - Send value to channel with syntax: channel_name <- value

	msg := <-messages //Receive value from channel with syntax: <- channel_name (main goroutine)
	fmt.Println(msg)
}
