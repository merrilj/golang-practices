package main

import "fmt"

type square struct {
	sideLength float64
}

type triange struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{4}
	t := triange{base: 6, height: 12}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}