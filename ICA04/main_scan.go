package main

import (
  "./scan"
  "fmt"
  "bufio"
  "os"
)

func main(){
  fmt.Println("Skriv inn filen du ønsker å skanne (files/...)")
  var filelength = bufio.NewScanner(os.Stdin)
  var filename string = "Velg en fil"
  for filelength.Scan() {
    filename = string(filelength.Text())
    break
  }
  scan.Skan(filename)
}