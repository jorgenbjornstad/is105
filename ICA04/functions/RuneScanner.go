package functions

import (
	"bufio"
	"os"
	"fmt"
	"is105/ICA04/fileutils"
)

func RuneScanner (){
fmt.Println("Skriv inn filen du ønsker å skanne (files/...)")
var filelength = bufio.NewScanner(os.Stdin)
var filename string = "Velg en fil"
for filelength.Scan() {
filename = string(filelength.Text())
break
}

fmt.Print("Antall linjeskift:")
fmt.Println(fileutils.SearchText(filename,"\n"))

m := fileutils.CountRune(filename)

fileutils.PrintRune(m)
}