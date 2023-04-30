package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var s = "สวัสดี"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
	fmt.Println()
	for idx , v := range s{
		fmt.Printf("%#U starts at %d\n", v, idx)
	}

	// Len: 18
	// e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 
	// Rune count: 6
	// U+0E2A 'ส' starts at 0
	// U+0E27 'ว' starts at 3
	// U+0E31 'ั' starts at 6
	// U+0E2A 'ส' starts at 9
	// U+0E14 'ด' starts at 12
	// U+0E35 'ี' starts at 15

	fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
	}
	// Using DecodeRuneInString
	// U+0E2A 'ส' starts at 0
	// U+0E27 'ว' starts at 3
	// U+0E31 'ั' starts at 6
	// U+0E2A 'ส' starts at 9
	// U+0E14 'ด' starts at 12
	// U+0E35 'ี' starts at 15
}