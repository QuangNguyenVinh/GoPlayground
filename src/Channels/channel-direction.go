package main

import "fmt"

//Only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
    pings <- msg
}
//Accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "PING PONG")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}