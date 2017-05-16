package main

import (
	"fmt"
	"strings"
	"./ascii"
	"unicode"
)

func main() {
	_, hellohex, _ := ascii.GreetingASCII()


	f := func(c rune) bool {
		return unicode.Is(unicode.ASCII_Hex_Digit, c)
	}
	fmt.Println(hellohex)
	if strings.IndexFunc(hellohex, f) != -1 {
		fmt.Println("All ASCII here")
	} else {
		fmt.Println("Found non-ASCII character")
	}
}
