package main

import "fmt"

func main(){

	for i := 0 ; i < 10; i++ {
		if(i % 2 == 0){
			fmt.Println("Even", i)
		} else {
			fmt.Println("Odd", i)
		}
	}

	if num:=-1; num == 0 {
		fmt.Println("Number is 0")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is positive")
	}
}