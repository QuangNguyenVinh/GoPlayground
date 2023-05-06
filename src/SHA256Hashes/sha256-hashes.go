package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Hello World"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s) // a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e
	fmt.Printf("%x\n", bs)
}
