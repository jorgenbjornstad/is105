package main

import ("github.com/nicolaifsf/go-speak"
	"fmt"
)

func main() {
	speech.SetWitKey("UMRVL4TFMAZEUL7U3GQ57ELPIOWSC5YS")
	fmt.Println(speech.SendWitVoice("wav/OSR_us_000_0010_8k.wav"))

}