package main

import "fmt"
import "./iso"

func main() {

  e := iso.ExtendedAscii
  h := iso.Hello

  var con bool

  for j := 0; j < len(h); j++ {
  for i := 0; i < len(e); i++ {
      if h[j] == e[i] {
        con = true
        break
      } else {
        con = false
      }
    }
    if con == false {
      fmt.Println("index", j, "is not part of extended ASCII")
    }
  }
}
