package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	arg := os.Args[1]
	fileInfo, err := os.Stat(arg)
	if err != nil {
		log.Fatal(err)
		}
		i := fileInfo.Size()

		varB := float64(i)
		varMB := varB / 1024
		varGB := varMB / 1024
		varTB := varGB / 1024
		fmt.Println("Information about a file", arg, ":")
		fmt.Println("Size:", varB, "in bytes.", varMB, "in megabytes.", varGB, " in gigabytes.", varTB, "in terrabytes.")
		fmt.Println("Is Directory: ", fileInfo.IsDir())
		fmt.Println("Is a regular file: ", fileInfo.Mode().IsRegular())
	        fmt.Println("Is append only: ", fileInfo.Mode()&os.ModeAppend)
		fmt.Println("Permissions:", fileInfo.Mode())
		fmt.Println("Last modified:", fileInfo.ModTime())
		fmt.Printf("System interface type: %T\n", fileInfo.Sys())
		fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	}

