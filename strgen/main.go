package main

import (
	"fmt"
	"math/rand"
	"time"
)

var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 18; i++ {
		idx := rand.Intn(len(charset))
		fmt.Printf("%c", charset[idx])
	}
	fmt.Println("")
}
