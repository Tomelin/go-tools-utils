package main

import (
	"fmt"

	"github.com/Tomelin/go-tools-utils/crypt"
)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func main() {

	var c crypt.Crypt
	StringToEncrypt := "Encrypting this string"
	// To encrypt the StringToEncrypt
	encText, err := c.Encrypt(StringToEncrypt)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	fmt.Println(encText)
	// To decrypt the original StringToEncrypt
	decText, err := c.Decrypt(encText)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)

	b64 := c.Base64Encrypt(StringToEncrypt)
	fmt.Println("ase 64 encrypt", b64)

	fmt.Println("Decrypt base64 ", c.Base64Decrypt(b64))
}
