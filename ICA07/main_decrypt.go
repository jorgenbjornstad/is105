package main

import "./decrypting"
import "fmt"

func main() {
decrypted := decrypting.Decrypt()
fmt.Println(decrypted)
}
