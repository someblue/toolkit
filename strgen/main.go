package main

import (
	"crypto/rand"
	"fmt"
)

var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	bs := make([]byte, 16)
	_, err := rand.Read(bs)
	if err != nil {
		panic(err)
	}
	for _, b := range bs {
		idx := int(b) % len(charset)
		fmt.Printf("%c", charset[idx])
	}
	fmt.Println("")
}
