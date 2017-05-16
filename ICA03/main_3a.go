package main

import "fmt"

func main() {

	fmt.Printf("Print med format %_s %s \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("Print med format %_q %q \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("Print med format %_+q %+q \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("Print med format %_c %c \n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

	fmt.Printf("%s \n", "\xc2\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
}
