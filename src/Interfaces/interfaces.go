package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

	//Empty Inteface
	var i interface{}

	describe(i) // (nil, nil)

	i = 02

	describe(i) // (2, int)

	i = "string"

	describe(i) // (string, string)

	person := make(map[string]interface{}, 0)

	person["name"] = "Quang"
	person["age"] = 18

	fmt.Println(person)

	//Type assertion
	var ix interface{} = "hello"

	s := ix.(string)
	fmt.Println(s)

	s, ok := ix.(string)
	fmt.Println(s, ok)

	f, ok := ix.(float64)
	fmt.Println(f, ok)

	f = ix.(float64) // Error
	fmt.Println(f)

}