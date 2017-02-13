package main

import "os"
import "fmt"
import "strconv"

func main() {

arg1 := os.Args[1]
arg2 := os.Args[2]

f1, err := strconv.ParseFloat(arg1, 64)
if err == nil {
  fmt.Println("  ",f1)
}else{
  fmt.Println("error")
}

f2, err2 := strconv.ParseFloat(arg2, 64)
if err2 == nil {
  fmt.Println("+ ",f2)
}else{
  fmt.Println("error")
}


sum = f1 + f2
fmt.Println("= ",sum)


}
