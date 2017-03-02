package main

import "fmt"
import "./iso"

func main() {

  var extascii []byte
	for i := 0x80; i <= 0xFF; i++ {
	extascii = append(extascii, byte(i))
	}

  h := iso.Hello

  var con bool
  for j := 0; j < len(h); j++ {
    for i := 0; i < len(extascii); i++ {
      if h[j] == extascii[i] {
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
