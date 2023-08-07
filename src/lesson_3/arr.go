package main

import "fmt"

func main() {
	slices()
}

func arr() {
	var a [10]int
	a[0] = 1
	fmt.Println(a)
	numbers := [...]int{1, 2, 3}
	fmt.Println(numbers)
}

func slices() {
	slices := []int{1, 2, 3}
	slices = append(slices, 4)
	fmt.Println(slices)

	slice1 := make([]string, 3)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1 = append(slice1, "e")
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	animals := [4]string{
		"cat",
		"dog",
		"mouse",
	}

	slice2 := animals[0:2]
	fmt.Println(slice2)
	slice2[0] = "chopi"
	slice2[1] = "serezha"
	fmt.Println(slice2)
	fmt.Println(animals)
	slice2 = append(slice2, "mua")
	slice2 = append(slice2, "mua")
	slice2 = append(slice2, "mua")
	slice2 = append(slice2, "mua")
	fmt.Println(slice2)
	fmt.Println(animals)
}
