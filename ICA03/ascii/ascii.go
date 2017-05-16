package ascii

import "fmt"

const Ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
    "\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
    ` !"#$%&'()*+,-./0123456789:;<=>?` +
    `@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
    "`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

func IterateOverASCIIStringLiteral() {
    for i := 0; i < len(Ascii); i++ {
        fmt.Printf("%b  %q  %x  \n", Ascii[i], Ascii[i], Ascii[i])
    }
}

func GreetingASCII() ([]byte, string, string) {

    hello := []byte {34, 72, 101, 108, 108, 111, 32, 58, 45, 41, 34}
  hellohex := "\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"

    konverterer := string(hello[:])
    return hello, hellohex, konverterer


}
