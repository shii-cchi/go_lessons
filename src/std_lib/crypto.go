package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Hello, world!")
	hash := sha256.Sum256(data)
	fmt.Printf("SHA-256 Hash: %x\n", hash)
}
