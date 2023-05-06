package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64) // 64 is bits of precision to parse.
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) //0 means infer the base from the string
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135") //a convenience function for basic base-10 int parsing
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
