package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Crypt struct {
	Keyphrase string `json:"keyphrase"`
	String    string `json:"string"`
	File      string `json:"file"`
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func (c *Crypt) encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func (c *Crypt) decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func (c *Crypt) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return c.encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func (c *Crypt) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := c.decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func (e *Crypt) Base64Encrypt(data string) string {
	return e.encode([]byte(data))
}

func (e *Crypt) Base64Decrypt(data string) string {
	cipherText := e.decode(data)
	return string(cipherText)
}
