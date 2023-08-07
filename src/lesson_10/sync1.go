package main

import (
	"sync"
	"time"
)

type counter1 struct {
	mu sync.RWMutex
	c  map[string]int
}

func (c *counter1) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *counter1) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c
}

func main() {
	key := "test"
	c := counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {

	}

	time.Sleep(time.Second * 3)
}
