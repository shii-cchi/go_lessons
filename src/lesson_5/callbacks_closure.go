package main

import "fmt"

func main() {
	//callbacks()
	closure()
}

func callbacks() {
	sumRes := calculate(sum, 5, 5)
	multRes := calculate(mult, 5, 5)
	fmt.Println(sumRes, multRes)
}

func sum(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func calculate(callback func(int, int) int, a, b int) int {
	result := callback(a, b)
	return result
}

func closure() {
	getTotalPrice := totalPrice(10)
	fmt.Println(getTotalPrice(100))
	fmt.Println(getTotalPrice(100))
	fmt.Println(getTotalPrice(100))
}

func totalPrice(initPrice int) func(int) int {
	return func(x int) int {
		initPrice += x
		return initPrice
	}
}
