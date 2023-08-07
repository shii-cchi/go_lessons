package main

import (
	exec "golang.org/x/sys/execabs"
	"os"
)

func main() {
	c := exec.Command("cat", "go_start/lesson_13/text.txt")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()

	if err != nil {
		panic(err)
	}
}
