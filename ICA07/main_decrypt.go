package main

import (
  "./decrypting"
  "fmt"
)

func main() {
decrypted := decrypting.Decrypt()
fmt.Println(decrypted)
}
