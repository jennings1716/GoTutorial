package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	sh string
	b  float64
	h  float64
}
type square struct {
	sh string
	a  float64
}

func main() {
	tri := triangle{b: 10, h: 10, sh: "Triangle"}
	printArea(tri)

	sq := square{a: 10, sh: "Square"}
	printArea(sq)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.b * t.h
}

func (s square) getArea() float64 {
	return s.a * s.a
}

func printArea(s shape) {
	fmt.Println("Area ", s.getArea())
}
