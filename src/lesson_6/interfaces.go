package main

import (
	"fmt"
	"math"
)

var _ Shape = &triangle{}
var _ Shape = &rectangle{}
var _ Shape = &circle{}

func main() {
	circle1 := circle{5}
	triangle1 := triangle{10, 2}
	rectangle1 := rectangle{5, 8}
	PrintSquare(circle1)
	PrintSquare(triangle1)
	PrintSquare(rectangle1)
}

type Shape interface {
	GetSquare() float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) GetSquare() float64 {
	return 0.5 * t.base * t.height
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) GetSquare() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) GetSquare() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func PrintSquare(s Shape) {
	fmt.Println(s.GetSquare())
}
