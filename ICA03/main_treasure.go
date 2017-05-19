package treasure

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func check(e error) {
	// Generic fatal error function
	if e != nil {
		log.Fatal(e)
	}
}

func DecodeTreasureFileUTF8(filename string) string {
	var byteslice []byte
	file, err := ioutil.ReadFile(filename)
	// Read a text file formatted like "\x20\x20..."
	check(err)
	for _, x := range strings.Split(string(file), "\\x") {
		// Split by \x and iterate over every value between
		if len(x) > 0 { // Make sure the string is not ""
			x = strings.TrimSpace(x)               // Remove any whitespace
			i, err := strconv.ParseUint(x, 16, 16) // Parse the string as base 16
			check(err)
			byteslice = append(byteslice, byte(i)) // Put it into byteslice
		}
	}
	reader := bytes.NewReader(byteslice) // Read byteslice
	// Decode from CP1252 to UTF8
	decoder := transform.NewReader(reader, charmap.Windows1252.NewDecoder())
	decoded_bytes, err := ioutil.ReadAll(decoder)
	check(err)
	return string(decoded_bytes) // Coerce []byte to string
}
