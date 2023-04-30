package main

import (
	"errors"
	"fmt"
)

func getMultiply(x, y, z int) int {
	return x * y * z
}

// Multiple return values
func getMaxValueFromMap(mapObject map[string]int) (int, error) {
	if len(mapObject) < 1 {
		return 0, errors.New("Map is empty")
	}
	arr := make([]int, len(mapObject))
	i := 0
	for _, v := range mapObject {
		arr[i] = v
		i += 1
	}
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max, nil
}

//Variadic functions
func getTotal(nums ...int) int {
	fmt.Print(nums, " ")
	sum := 0
	for _, v := range nums{
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Multiply:", getMultiply(1, 2, 3))

	mapObject := map[string]int{"G": 1, "O": 3, "L": -5, "A": 7, "N": -3}

	maxValue, message := getMaxValueFromMap(mapObject)
	if message == nil {
		fmt.Println("Max value from map is", maxValue)
	} else {
		fmt.Println(message)
	}

	fmt.Println("Total:", getTotal(1, 2, 3, 4, 5))

}
