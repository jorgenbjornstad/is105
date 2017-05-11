package main

import "./crypting"
import "fmt"

func main() {
encrypted := crypting.Encrypt("snart sommerferie")
fmt.Println(encrypted)
}
