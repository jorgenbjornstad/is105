package main

import "fmt"
import "./ascii"


func main() {
a := ascii.Ascii
h := ascii.Hello

var con bool

for j := 0; j < len(h); j++ {
for i := 0; i < len(a); i++ {
    if h[j] == a[i] {
      con = true
      break
    } else {
      con = false
    }
  }
    if con == false {
      fmt.Println("index", j, "is not part of ASCII")
    }
    }
  }
}
