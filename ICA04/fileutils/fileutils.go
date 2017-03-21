package fileutils

import (
	"io"
	"log"
	"os"
	"bufio"
	"fmt"
)

func FileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	return byteSlice
}


func BReader (filename string) (*bufio.Reader) {
	// Open file and create a buffered reader on top
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewReader(f)
}

func NumberOfLines(filename string) {
	// Loads the buffered file
	reader := BReader(filename)
	// Read number of lines until end of file
	for {
		line, err := reader.ReadString(0x0A)
		fmt.Printf("%+q\n", line) // Display newlines with %+q
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
