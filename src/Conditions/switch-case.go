package main

import (
	"fmt"
	"time"
)

func main(){

	i := 0

	switch i {
		case 0:
			fmt.Println("Zero")
		case -1:
			fmt.Println("Negative one")
		case 1:
			fmt.Println("Positive one")
	}

	switch time.Now().Weekday(){ 
		case time.Saturday, time.Sunday:
			fmt.Println("Weekend")
		default:
			fmt.Println("Working day")
	}

	t := time.Now()
	switch{
		case t.Hour() < 12:
			fmt.Println("AM")
		default:
			fmt.Println("PM")
	}

	checkType := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default: 
		fmt.Printf("Haven't implemented type \"%T\" yet!", t)
		}
	}

	checkType("string")

}