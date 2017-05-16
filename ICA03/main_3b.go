package main

import (
  "./fileutils"
  "fmt"

)

func main () {
    // Constructs a byteslice from file.
    i := fileutils.FileToByteslice("files/lang01.wl")
    e := fileutils.FileToByteslice("files/lang02.wl")
    f := fileutils.FileToByteslice("files/lang03.wl")

    // Prints each constructed slice with %X.
    fmt.Printf("% X \n", i[:])
    fmt.Printf("% X \n", e[:])
    fmt.Printf("% X \n", f[:])
}