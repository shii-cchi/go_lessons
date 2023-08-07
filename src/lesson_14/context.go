package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создание корневого контекста
	ctx := context.Background()

	// Создание дочернего контекста с возможностью отмены
	ctx, cancel := context.WithCancel(ctx)

	// Запуск горутины для выполнения операции
	go performTask(ctx)

	// Ждем 3 секунды, затем отменяем выполнение операции
	time.Sleep(3 * time.Second)
	cancel()

	// Ожидаем завершения горутины
	time.Sleep(1 * time.Second)

	fmt.Println("Главная горутина завершена")
}

func performTask(ctx context.Context) {
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
