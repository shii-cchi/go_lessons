package main

import (
	"errors"
	"fmt"
)

func main() {
	var a float64 = 32
	var b float64 = 2
	result, err := Divide(a, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат =", result)
	}
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}
