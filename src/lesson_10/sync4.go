package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	done := make(chan bool)

	for i := 0; i < 30; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("only one")
			})
			done <- true
		}()
	}

	for i := 0; i < 30; i++ {
		<-done
	}

}
