package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p) // printing “verbs” designed to format general Go values

	fmt.Printf("struct2: %+v\n", p) // will include the struct’s field names.

	fmt.Printf("struct3: %#v\n", p) // representation of the value, i.e. the source code snippet that would produce that value.

	fmt.Printf("type: %T\n", p) // the type of a value

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 14) // a binary representation.

	fmt.Printf("char: %c\n", 33) // prints the character corresponding to the given integer.

	fmt.Printf("hex: %x\n", 456) // provides hex encoding.

	fmt.Printf("float1: %f\n", 78.9)
	//%e and %E format the float in (slightly different versions of) scientific notation.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"") //string

	fmt.Printf("str2: %q\n", "\"string\"") // double-quote strings

	fmt.Printf("str3: %x\n", "hex this") // renders the string in base-16, with two output characters per byte of input.

	fmt.Printf("pointer: %p\n", &p) // a representation of a pointer

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // the width of an integer

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // left-justify with - flag

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

	/*
		Output
		struct1: {1 2}
		struct2: {x:1 y:2}
		struct3: main.point{x:1, y:2}
		type: main.point
		bool: true
		int: 123
		bin: 1110
		char: !
		hex: 1c8
		float1: 78.900000
		float2: 1.234000e+08
		float3: 1.234000E+08
		str1: "string"
		str2: "\"string\""
		str3: 6865782074686973
		pointer: 0xc0000a6070
		width1: |    12|   345|
		width2: |  1.20|  3.45|
		width3: |1.20  |3.45  |
		width4: |   foo|     b|
		width5: |foo   |b     |
		sprintf: a string
		io: an error

	*/
}
