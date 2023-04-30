package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

func (c *circle) area() float64 {
	return float64(c.radius) * float64(c.radius) * math.Pi
}

func (c circle) perim() float64 {
	return 2 * float64(c.radius) * math.Pi
}

func main() {
	c := circle{radius: 10}

	fmt.Println("area: ", c.area())
	fmt.Println("perim:", c.perim())

	cp := &c
	fmt.Println("area: ", cp.area())
	fmt.Println("perim:", cp.perim())
}
