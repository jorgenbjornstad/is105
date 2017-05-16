package main

import (
	"fmt"
	"os"
	"strconv"

	"./functions"
)

//bruker log funksjonen for å regne ut log2 av gitt tall i float64
//Tar parameter fra komandolinje og gjør det om til float64
func main() {
	args1 := os.Args[1]
	param, _ := strconv.ParseFloat(args1, 64)
	fmt.Println(functions.Log2(param))
}
