package main

import (
	"encoding/hex"
	"goCipher/gcm"
	"log"
)

func main() {
	BaseKey := "123456789"
	plainText := "a123"
	c := gcm.NewCipher(BaseKey)
	cipherString, err := c.Encrypt([]byte(plainText))
	if err != nil {
		log.Fatal(err)
	}
	cipherText := hex.EncodeToString(cipherString)
	log.Println("cipherText=", cipherText)

	v, _ := hex.DecodeString(cipherText)
	rawText, err := c.Decrypt(v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("rawText=", string(rawText))
}
