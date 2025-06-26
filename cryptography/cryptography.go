package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/joho/godotenv"
	"github.com/legiorex/manager-password/output"
)

type Cryptography struct {
	Key string
}

func NewCryptography() *Cryptography {
	myEnv, err := godotenv.Read()

	if err != nil {
		output.PrintError("Не удалось загрузить evn файл")
	}

	key := myEnv["KEY"]

	if key == "" {
		panic("Не передан параметр KEY в env файле")
	}

	return &Cryptography{
		Key: key,
	}
}

func (cry *Cryptography) Encrypt(encryptBytes []byte) []byte {

	block, err := aes.NewCipher([]byte(cry.Key))

	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)

	if err != nil {
		panic(err.Error())
	}

	return aesGCM.Seal(nonce, nonce, encryptBytes, nil)

}

func (cry *Cryptography) Decrypt(decryptBytes []byte) []byte {
	block, err := aes.NewCipher([]byte(cry.Key))

	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, cipherDecrypt := decryptBytes[:nonceSize], decryptBytes[nonceSize:]

	result, err := aesGCM.Open(nil, nonce, cipherDecrypt, nil)

	if err != nil {
		panic(err.Error())
	}

	return result

}
