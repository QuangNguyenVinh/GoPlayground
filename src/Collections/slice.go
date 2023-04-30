package main

import "fmt"

func main() {

	s := make([]string, 3)

	s[0] = "G"
	s[1] = "O"
	s[2] = "-"

	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "L")
	s = append(s, "A", "N", "G")

	fmt.Println(s) // G O - L A N G

	//Copy slice
	s2 := make([]string, len(s))
	copy(s2, s) // Copy all element from s to s2
	fmt.Println(s2)

	fmt.Println(s2[:2]) // G O

	t := []string{"G", "O", "!"}
	fmt.Println("Init:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := (i + 1) * 2
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // [0 1] [1 2 3 4] [2 3 4 5 6 7]]
}
