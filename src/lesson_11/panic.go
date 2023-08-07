package main

import (
	"fmt"
	"log"
)

func main() {
	//panic("something goes wrong")

	divide(1, 0)
}

func divide(a, b int) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic happened", err)
		}
	}()

	fmt.Println(a / b)
}
