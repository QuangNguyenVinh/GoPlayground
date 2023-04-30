package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    for index, num := range nums {
		fmt.Printf("Index %d and value %d\n", index, num)
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}