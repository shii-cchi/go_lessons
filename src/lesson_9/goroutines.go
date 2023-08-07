package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go say("hello", ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func say(word string, ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
