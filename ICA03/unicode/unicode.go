package unicode

import "os"
import "fmt"

// Kode for Oppgave 4a
func Translate() {

	x := "nor\u00F0ur og su\u00F0ur"
	y := "\uF963\u5357\u30C8"

	a := "is"
	b := "jp"
	c := "nord og sør"

	var con bool
	con = false

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	if (arg1 == c) {
		con = true
		if arg2 == a {
			fmt.Printf("%q \n", x)
		} else if arg2 == b {
			fmt.Printf("%q \n", y)
	} else {
		fmt.Println("Skrev du inn riktig språk?")
		fmt.Println("Skriv inn \"is\" eller \"jp\" for å oversette til islands eller japansk")
	}
}
	if con == false {
		fmt.Println("Tilgjengelige uttytkk å oversette:")
		fmt.Println("\"nord og sør\"")
	}
}
