package decrypting

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Decrypt() string {

	message, err := ioutil.ReadFile("C:/Work/src/is105/ICA07/message/message.txt")

	key := []byte("ellevilleellever")

	var iv = []byte("abcdef1234567890")[:aes.BlockSize]

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(1)
	}
	fmt.Printf("NewCipher(%d bytes)\n", len(key))


	decrypter := cipher.NewCFBDecrypter(block, iv)

	decrypted := make([]byte, len(message))
	decrypter.XORKeyStream(decrypted, message)
	// fmt.Printf("Decrypting %v -> %v : %s\n", message, decrypted, decrypted)
	decryptedfinal := string(decrypted[:])
	return decryptedfinal
}
