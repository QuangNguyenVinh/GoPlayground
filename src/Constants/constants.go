package main

import (
	"fmt"
	"math"
)
const str string = "Hello World"

func main(){
	fmt.Println(str)

	const pi float64 = 3.14
	fmt.Println(pi)
	fmt.Println(int64(pi)) //Can't convert constant
	fmt.Println(math.Abs(-2))
}