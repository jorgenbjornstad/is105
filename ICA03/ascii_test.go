package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(r rune) bool {
		return r < 'A' || r > 'z'
	}
	if strings.IndexFunc("HelloWorld", f) != -1 {
		fmt.Println("Found special char")
	}
	if strings.IndexFunc("Hello World", f) != -1 {
		fmt.Println("Found special char")
	}
}
