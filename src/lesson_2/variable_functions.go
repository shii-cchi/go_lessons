package main

import "fmt"

func main() {
	const a string = "name"

	fmt.Println(a)
	fmt.Println(test())
	fmt.Println(test_1())

	s1, s2 := test()
	fmt.Println(s1, s2)

	fmt.Println(test_2())
}

func test() (string, string) {
	a := "hello"
	b := "world"
	return a, b
}

func test_1() string {
	a := "hello"
	return a
}

func test_2() (a, b, c string) {
	a = "hello"
	b = "world"
	c = "f"
	return a, b, c
}
