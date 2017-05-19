package main

import "./huffman"
import "fmt"

/*
a := 000
b := 001
c := 111
d := 10
e := 110
f := 01
*/

func main(){
  x, _ := huffman.Decode("1110000011111011001")
  fmt.Println(x)
}
