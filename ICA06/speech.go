package main

import ("github.com/nicolaifsf/go-speak"
	"fmt"
)

func main() {
	speech.SetWitKey("JPTUX4UZQXYNBIZCGRRCB6GKPNRVJFUN")
	fmt.Println(speech.SendWitVoice("wav/vebjorn_engelsk.wav"))
	fmt.Println(speech.SendWitVoice("wav/snorre_engelsk.wav"))

}
