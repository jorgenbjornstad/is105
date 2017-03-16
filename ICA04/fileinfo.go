package main

import (
	"fmt"
	"log"
	"os"
	"./files"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {

	fileInfo, err = os.Stat("C:/Work/src/is105/ICA04/files/pg100.txt)
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
