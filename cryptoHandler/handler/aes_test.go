package handler

import (
	"log"
	"testing"
)

func TestAscii(t *testing.T) {
	i := 0
	log.Println(rune(i))
	log.Println(string(rune(i)))
	log.Println(len([]byte(string(rune(i)))))
}
func TestEncrypt(t *testing.T) {
	key := "0123456789876543"
	text := "hello world!"
	aes := AesHandler{}
	cypherText, err := aes.Encrypt(key, text)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cypherText)
	originText, err := aes.Decrypt(key, cypherText)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(originText)

}
