package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	expression := "2 + 2, 3 * 4, 6 / 0, 5 - 2"
	operationsArr := strings.Split(expression, ", ")

	resultChan := make(chan int)
	errorChan := make(chan error)
	mapResults := make(map[string]int)

	for _, operation := range operationsArr {
		go calculate(operation, resultChan, errorChan)
		err := <-errorChan

		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			<-resultChan
		} else {
			mapResults[operation] = <-resultChan
		}
	}

	printMap(mapResults)
}

func calculate(operation string, resultChan chan int, errorChan chan error) {
	operator, operand1, operand2 := splitExpression(operation)

	var result int
	var err error

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			err = errors.New("деление на ноль")
		} else {
			result = operand1 / operand2
		}
	}

	errorChan <- err
	resultChan <- result
}

func splitExpression(operation string) (string, int, int) {
	expressionArr := strings.Split(operation, " ")
	operator := expressionArr[1]
	operand1, _ := strconv.Atoi(expressionArr[0])
	operand2, _ := strconv.Atoi(expressionArr[2])

	return operator, operand1, operand2
}

func printMap(mapResults map[string]int) {
	for operation, result := range mapResults {
		fmt.Printf("%s = %d\n", operation, result)
	}
}
