package main

import (
	"fmt"
	"strings"
	"./ascii"
	"unicode"
)

func main() {
	hello, _, _ := ascii.GreetingASCII()


	f := func(c rune) bool {
		return unicode.Is(unicode.ASCII_Hex_Digit, c)
	}

	if strings.IndexFunc(hello, f) != -1 {
		fmt.Println("Found non-ASCII character")
	} else {
		fmt.Println("All ASCII here")
	}
}
