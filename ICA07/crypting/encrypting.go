package crypting

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Encrypt(message string) {

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

	os.Remove("C:/Work/src/is105/ICA07/message/message.txt")
	newfile, err := os.Create("C:/Work/src/is105/ICA07/message/message.txt")
  check(err)
	defer newfile.Close()

	newfile.Write(encrypted)
	fmt.Println("Message encrypted.")
}
