package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	fmt.Println("Sorted Ints: ", sort.IntsAreSorted(ints))
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Sorted Ints Reverse: ", ints)
	fmt.Println("Sorted Strings: ", sort.StringsAreSorted(strs))
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Println("Sorted Strings Reverse: ", strs)
}
