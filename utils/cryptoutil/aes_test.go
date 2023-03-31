package cryptoutil

import (
	"encoding/hex"
	"log"
	"strconv"
	"strings"
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

func TestDecryptECBFromJava(t *testing.T) {
	ks := "[-9, -126, 104, -32, -65, 80, -2, 39, -23, 48, 63, 20, -90, 27, 44, -64]"
	if strings.Index(ks, "[") == 0 {
		ksr := ks[1 : len(ks)-1]
		log.Printf("ksr:%v", ksr)
		sa := strings.Split(ksr, ",")
		var keys = make([]byte, len(sa))

		for i := 0; i < len(sa); i++ {
			n, error := strconv.Atoi(strings.TrimSpace(sa[i]))
			if error != nil {
				log.Fatal(error)
			}
			//上述字节来自java,而golang中byte(0-255),负数+256
			keys[i] = byte(n)
		}

		text := "CFA3F485507F35F53EF4CF08FEF3BDB2E49D958DD91A61672E6DD92C5948FDD3"

		aes := AES{}
		encrypted, err := hex.DecodeString(text)
		if err != nil {
			panic(err)
		}
		decrypted, err := aes.AESDecryptECBImpl(keys, encrypted)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ksr:%v", string(decrypted))
	}

}
