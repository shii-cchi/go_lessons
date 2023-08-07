package main

import "fmt"

func main() {

	var name = "Chopi"
	var name2 string
	name1 := "CHOPICK"

	var (
		name5 = "flop"
		age   = 111
	)

	fmt.Println(name2)
	fmt.Println(name1)
	fmt.Println(name5)

	d := fmt.Sprintf("My name is %s. And my age is %d", name, age)
	fmt.Println(d)
}
