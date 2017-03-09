package main

import (
"is105/ICA03/fileutils"
	"fmt"


)
func main() {
	byteslice1 := fileutils.FileToByteslice("files/text1.txt")
	byteslice2 := fileutils.FileToByteslice("files/text2.txt")
	// fmt.Println(byteslice1)
	// fmt.Printf("%s", byteslice1)
	// fm.Printf("%q", byteslice1
        fmt.Println(byteslice1)
	fmt.Printf("% s", byteslice1)

	fmt.Println(byteslice2)
	fmt.Printf("% s", byteslice2)

}
