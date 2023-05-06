package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))     // sec & nsec
	fmt.Println(time.Unix(0, now.UnixNano())) //sec & nsec

	/*
		Output:
		2023-05-06 14:53:00.6905245 +0700 +07 m=+0.001029101
		1683359580
		1683359580690
		1683359580690524500
		2023-05-06 14:53:00 +0700 +07
		2023-05-06 14:53:00.6905245 +0700 +07
	*/
}
