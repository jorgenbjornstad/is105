package main

import "os"
import "fmt"
import "strconv"

func main() {

arg1 := os.Args[1]
arg2 := os.Args[2]


f1, err := strconv.ParseUint(arg1, 10, 32)
if err == nil {
  fmt.Println("  ",f1)
}else{
  fmt.Println("error")
}


f2, err2 := strconv.ParseUint(arg2, 10, 32)
if err2 == nil {
  fmt.Println("+ ",f2)
}else{
  fmt.Println("error")
}

// Sum er int64, ikke int32.
// Prøvd å få sum til å være noe annet enn int64, men får det ikke til.
sum := f1 + f2
fmt.Println("= ",sum)
}
