package main

import "fmt"

func main() {
	var a float64 = 55
	var b float64 = 0
	result, err := DivideFloat(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

type CustomError struct {
	Code int
	Msg  string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", c.Code, c.Msg)
}

func DivideFloat(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &CustomError{Code: 1, Msg: "divide by zero"}
	}
	return a / b, nil
}
