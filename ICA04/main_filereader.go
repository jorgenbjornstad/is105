package main

import (
	"./fileutils"
)

func main() {
	fileutils.BReader("files/text1.txt")
	fileutils.NumberOfLines("files/text1.txt")
}