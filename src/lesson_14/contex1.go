package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	duration := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, duration)

	go performTask1(ctx)

	defer cancel()

	time.Sleep(6 * time.Second)
}

func performTask1(ctx context.Context) {
	// Симуляция длительной операции
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Выполняется операция...")
		case <-ctx.Done():
			// Операция была отменена
			fmt.Println("Операция отменена")
			return
		}
	}
}
