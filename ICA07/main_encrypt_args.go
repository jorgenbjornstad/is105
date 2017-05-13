package main

import (
  "./crypting"
  "os"
)

func main() {
  arg := os.Args[1]
  crypting.Encrypt(arg)
}
