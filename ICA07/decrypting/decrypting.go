package decrypting

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func Decrypt(message []byte) string {

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
