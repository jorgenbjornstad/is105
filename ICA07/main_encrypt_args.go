package main

import "./crypting"
import "os"

func main() {
  arg := os.Args[1]
  crypting.Encrypt(arg)
}
