package main

import "./decrypting"
import "fmt"

// Vil endre denne og main_encrypt.go til å kunne åpne en fil og bruke innholdet i filen
// som parameter slik at vi kan skrive en hemmelig melding i en tekstfil for så å kunne
// kryptere den og få en slice av bytes i en tesktfil som vi kan dekryptere.
func main() {
var buf = []byte{168, 22, 110, 191, 13, 173, 13, 84, 5, 94, 213, 54, 91, 156, 15, 177, 246}
decrypted := decrypting.Decrypt(buf)
fmt.Println(decrypted)
}
