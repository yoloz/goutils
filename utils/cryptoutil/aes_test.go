package cryptoutil

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
func TestCBC(t *testing.T) {
	key := "0123456789876543"
	text := "hello world!"
	log.Printf("origin:%s", text)
	aes := AES{}
	cypherText, err := aes.AESEncryptCBC(key, text)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("encrypt:%s", cypherText)
	originText, err := aes.AESDecryptCBC(key, cypherText)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("decrypt:%s", originText)

}

func TestECB(t *testing.T) {
	key := "0123456789876543"
	text := "hello world!"
	log.Printf("origin:%s", text)
	aes := AES{}
	cypherText, err := aes.AESEncryptECB(key, text)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("encrypt:%s", cypherText)
	originText, err := aes.AESDecryptECB(key, cypherText)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("decrypt:%s", originText)

}

func TestCFB(t *testing.T) {
	key := "0123456789876543"
	text := "hello world!"
	log.Printf("origin:%s", text)
	aes := AES{}
	cypherText, err := aes.AESEncryptCBC(key, text)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("encrypt:%s", cypherText)
	originText, err := aes.AESDecryptCBC(key, cypherText)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("decrypt:%s", originText)

}
