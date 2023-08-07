package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var counter11 uint64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		k := i
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter11, 1)
			}

			fmt.Printf("%d\n", k)
			time.Sleep(300 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Printf("all done, couner = %d\n", counter11)

}
