package iso

import "fmt"

const Ascii = `€ ‚ƒ„…†‡ˆ‰Š‹Œ Ž ` +
	` ‘’“”•–—˜™š›œ žŸ` +
	` ¡¢£¤¥¦§¨©ª«¬¬®¯` +
	`°±²³´µ¶•¸¹º»¼½¾¿` +
	`ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ` +
	`ÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞß` +
	`àáâãäåæçèéêëìíîï` +
	`ðñòóôõö÷øùúûüýþÿ`

func IterateOverExtendedASCIIStringLiteral(sl string) {
	for i := 0; i < len(Ascii); i++ {
		fmt.Printf("%b  %q  %x  \n", Ascii[i], Ascii[i], Ascii[i])
	}
}

// Kode for Oppgave 2b
func GreetingExtendedASCII() {}
