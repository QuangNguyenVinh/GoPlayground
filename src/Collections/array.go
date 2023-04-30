package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println("Array:", arr)       // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("Length:", len(arr)) // 10

	init_arr := [5]int{10, 5, 3, 15, 7}
	fmt.Println(init_arr)

	var two_darr [3][4]int
	for i := 0; i < 3; i += 1 {
		for j := 0; j < 4; j += 1 {
			two_darr[i][j] = i * j
		}
	}
	fmt.Println("2d-arr", two_darr) //[0 0 0 0] [0 1 2 3] [0 2 4 6]]
}
