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
	varKiB := varB / 1024
	varMiB := varKiB / 1024
	varGiB := varMiB / 1024

	fmt.Println("Information about a file", arg, ":")
	fmt.Println("Size:", varB, "in bytes.", varKiB, "in kibibytes.", varMiB,
		" in mibibytes.", varGiB, "in gibibytes.")
	fmt.Println("Is Directory: ", fileInfo.Mode().IsDir())
	fmt.Println("Is a regular file: ", fileInfo.Mode().IsRegular())
	fmt.Println("Has UNIX permissions in bits: ", fileInfo.Mode())
	fmt.Println("Is append only: ", fileInfo.Mode()&os.ModeAppend != 0)
	fmt.Println("Is a device file:", fileInfo.Mode()&os.ModeDevice != 0)
	fmt.Println("Is a UNIX character device: ", (fileInfo.Mode()&os.ModeDevice != 0)&&(fileInfo.Mode()&os.ModeCharDevice != 0))
	fmt.Println("Is a UNIX block device: ", (fileInfo.Mode()&os.ModeDevice != 0)&&(fileInfo.Mode()&os.ModeCharDevice == 0))
	fmt.Println("Is a symbolic link: ", fileInfo.Mode()&os.ModeSymlink != 0)

}

