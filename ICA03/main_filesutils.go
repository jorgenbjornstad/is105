package main

import (
	"fmt"

	"./fileutils"
)

func main() {
	i := fileutils.FileToByteslice("C:/work/src/is105/ICA03/files/lang01.wl")
	e := fileutils.FileToByteslice("C:/work/src/is105/ICA03/files/lang02.wl")
	f := fileutils.FileToByteslice("C:/work/src/is105/ICA03/files/lang03.wl")
	//fileTobyteslice.FileToByteslice(lang01.go, string) []
	fmt.Printf("% X\n", i)
	fmt.Printf("% X\n", e)
	fmt.Printf("% X\n", f)
}
