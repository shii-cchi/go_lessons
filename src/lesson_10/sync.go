package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	key := "test"
	c := counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(c.Value(key))
}
