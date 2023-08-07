package main

import "fmt"

type Points struct {
	X int
	Y int
}

func movePoint(point *Points, shiftX, shiftY int) {
	point.X += shiftX
	point.Y += shiftY
}

func (point *Points) movingPoint(shiftX, shiftY int) {
	point.X += shiftX
	point.Y += shiftY
}

func main() {
	point := Points{1, 2}
	fmt.Println(point)

	movePoint(&point, 2, 2)
	fmt.Println(point)

	point.movingPoint(100, -100)
	fmt.Println(point)
}
