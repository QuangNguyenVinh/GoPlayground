package main

import "fmt"

func main() {

    m := make(map[string]int) // To create an empty map, use the builtin make: make(map[key-type]val-type).

    m["g"] = 7
    m["o"] = 13

    fmt.Println("Map:", m)

    v1 := m["g"]
    fmt.Println("Value for key g:", v1)

    v3 := m["o"]
    fmt.Println("Value for key o:", v3)

    fmt.Println("Length of map:", len(m))

    delete(m, "lang")
    fmt.Println("Map after delete key lang:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs) // If the key doesn’t exist, the zero value of the value type is returned.  The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("Map initial:", n)
}