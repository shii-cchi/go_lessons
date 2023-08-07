package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")

	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf(" n = %d err = %v arr = %v\n", n, err, arr)
		fmt.Println("n bytes:", arr[:n])
		if err == io.EOF {
			break
		}
	}
}
