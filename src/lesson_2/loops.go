package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
		fmt.Println(sum)
	}

	if sum > 1000 {
		fmt.Println("1)sum > 1000")
	} else {
		fmt.Println("1)sum < 1000")
	}

	for sum < 1000 {
		sum += 100
		switch sum {
		case 210:
			fmt.Println(sum + 5)
		case 610:
			fmt.Println(sum + 54)
		default:
			fmt.Println("another")
		}
	}

	if sum > 1000 {
		fmt.Println("2)sum > 1000")
	} else {
		fmt.Println("1)sum < 1000")
	}

	a := 2
	fmt.Println(isTest(a))
}

func isTest(a int) int {
	if a == 2 {
		return 2
	}
	return 3
}
