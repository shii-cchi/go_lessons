package main

import "fmt"

type Points struct {
	X int
	Y int
}

func (p Points) method() {
	fmt.Println(p.Y)
	fmt.Println(p.X)
}

func main() {
	myPoint := Points{
		X: 3,
		Y: 4,
	}

	myPoint.method()
}

func structs() {
	point := Points{
		X: 1,
		Y: 2,
	}
	fmt.Println(point)
	point.X = 3
	fmt.Println(point.X)
	fmt.Println(point.Y)

	point1 := Points{
		X: 111,
	}
	fmt.Println(point1)
}
