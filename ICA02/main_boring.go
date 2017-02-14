package main

import (
	"fmt"

	"./boring"
)

func main() {
	go boring.Boring01("Boring")
	for i := 0; ; i++ {
		fmt.Println("You're boring; I'm leaving.")
	}
}
