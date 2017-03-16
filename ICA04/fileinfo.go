package main

import (
	"fmt"
	"log"
	"os"
	"./files"
	"io/ioutil"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
        arg := os.Args[1]
	fileInfo, err = os.Stat(arg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
