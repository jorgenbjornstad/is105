package main

import ("github.com/nicolaifsf/go-speak"
	"fmt"
)

func main() {
	speech.SetWitKey("UMRVL4TFMAZEUL7U3GQ57ELPIOWSC5YS")
	fmt.Println(speech.SendWitVoice("wav/vebjorn_engelsk.wav"))

}