package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("go_start/lesson_13/text.txt")

	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}

	addData := []byte("BBBBBBBBBBBBBBBBB")
	err1 := os.WriteFile("go_start/lesson_13/text.txt", addData, 0777)

	if err1 != nil {
		panic(err1)
	}

	data, err = os.ReadFile("go_start/lesson_13/text.txt")

	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}

	f, err2 := os.OpenFile("go_start/lesson_13/text.txt", os.O_APPEND|os.O_WRONLY, 0777)
	defer f.Close()
	if err2 != nil {
		panic(err2)
	}

	if _, err3 := f.WriteString("myau"); err3 != nil {
		panic(err3)
	}

	data, err = os.ReadFile("go_start/lesson_13/text.txt")

	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}
}
