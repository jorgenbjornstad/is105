package huffman

import (
  "fmt"
)


func Decode(code string) {


  var u string
  maxbits := 8000

  for i := 0; i < maxbits; {
    if (code[i] == 0){
      i++
      if (code[i] == 0) {
        i++
        if (code[i] == 0) {
          u = "A"
          fmt.Print("\n", u)
          i++
        } else if (code[i] == 1) {
          u = "B"
          fmt.Print("\n", u)
          i++
        }
        } else if (code[i] == 1) {
          u = "F"
          fmt.Print("\n", u)
          i++
        }
      } else if (code[i] == 1) {
        i++
        if (code[i] == 0) {
          u = "D"
          fmt.Print("\n", u)
          i++
        } else if (code[i] == 1) {
          i++
          if (code[i] == 0) {
            u = "E"
            fmt.Print("\n", u)
            i++
          } else if (code[i] == 1) {
            u = "C"
            fmt.Print("\n", u)
            i++
          }
      }
    }
  }
}
