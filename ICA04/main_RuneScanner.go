package main

import (
	"./fileutils"
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Skriv inn fil (files/...)")
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