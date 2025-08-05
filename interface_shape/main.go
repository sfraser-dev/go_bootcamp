package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	// fields / properties
	base   float64
	height float64
}

type square struct {
	// fields / properties
	sideLength float64
}

func main() {
	t := triangle{
		base:   10,
		height: 2,
	}

	s := square{
		sideLength: 3,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := math.Pow(s.sideLength, 2)
	return area
}

func printArea(sh shape) {
	a := sh.getArea()
	fmt.Println("Area of shape is:", a)
}
