package main

import "fmt"
import "./ascii"


func main() {
a := ascii.Ascii
h, _, _ := ascii.GreetingASCII()

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
    } else {
      fmt.Println("index", j, "is a part of ASCII")
    }
  }
}
