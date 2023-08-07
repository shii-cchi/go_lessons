package main

import "fmt"

func main() {
	var a interface{}
	a = "hello"
	fmt.Printf("%v %T\n", a, a)
	a = 64
	fmt.Printf("%v %T\n", a, a)

	var b interface{} = "hello"
	s, ok := b.(string)
	fmt.Println(s, ok)
	d, ok := b.(int)
	fmt.Println(d, ok)
	//g := b.(int) panic
	//fmt.Println(g)

	checkType(b)
	checkType(a)
}

func checkType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case float64:
		fmt.Println("type is float64")
	default:
		fmt.Println("another interface type")
	}
}
