package main

import "fmt"

func main() {

	fmt.Printf("Print med format %%s %s \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("Print med format %%q %q \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("Print med format %%+q %+q \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("Print med format %%c %c \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

	fmt.Printf("La til en ekstra karakter (xc2) for å få 1/2: %s \n", "\xc2\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
}
