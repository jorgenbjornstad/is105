package crypting

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func Encrypt(message string) []byte {

	src := []byte(message)

	key := []byte("ellevilleellever")

	var iv = []byte("abcdef1234567890")[:aes.BlockSize]

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(1)
	}
	fmt.Printf("NewCipher(%d bytes)\n", len(key))

	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(src))
	encrypter.XORKeyStream(encrypted, src)
	// fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)
	return encrypted
}
