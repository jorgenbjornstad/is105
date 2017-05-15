package fileutils

import (
	"io"
	"log"
	"os"
	"bytes"
	"fmt"
	"sort"
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

func SearchText(filename string, detect string) int {
	// Load the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Buffer set to handle large files
	buffer := make([]byte, 32*4096)
	count := 0
	// Convert string to byte in order to search
		search := []byte(detect)
		c, _ := file.Read(buffer)
		count += bytes.Count(buffer[:c], search)
	return count
}

func CountRune(filename string) map[int]string {
	m := make(map[int]string)
	count := 0
	for i := 0x20; i <= 0x7F; i++ {
		count = SearchText(filename, string(i))
		rune := string(i)
		m[count] = rune
	}
	return m
}

func PrintRune(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

// Prints the five most read runes in the file.
	fmt.Println("5 mest forekomne runer:")
	rune1 := keys[len(keys)-1]
	rune2 := keys[len(keys)-2]
	rune3 := keys[len(keys)-3]
	rune4 := keys[len(keys)-4]
	rune5 := keys[len(keys)-5]
	fmt.Println("Antall:", rune1, "Rune:", m[rune1])
	fmt.Println("Antall:", rune2, "Rune:", m[rune2])
	fmt.Println("Antall:", rune3, "Rune:", m[rune3])
	fmt.Println("Antall:", rune4, "Rune:", m[rune4])
	fmt.Println("Antall:", rune5, "Rune:", m[rune5])
}

