package main

import (
  "./fileutils"
  "fmt"

)

func main () {
  i := fileutils.FileToByteslice("files/lang01.wl")
  e := fileutils.FileToByteslice("files/lang02.wl")
  f := fileutils.FileToByteslice("files/lang03.wl")
  //fileTobyteslice.FileToByteslice(lang01.go, string) []
  fmt.Printf("% X \n", i[:])
  fmt.Printf("% X \n", e[:])
  fmt.Printf("% X \n", f[:])
}
