package main

import "./huffman"
import "fmt"

func main(){
  print, _ := huffman.DecodeHuffmanString("1110000011111011001")
  fmt.Println(print)
}
