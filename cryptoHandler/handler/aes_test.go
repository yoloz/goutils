package handler

import (
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
func TestEncrypt(t *testing.T) {
	key := "0123456789876543"
	text := "hello world!"
	log.Printf("origin:%s", text)
	aes := AesHandler{}
	cypherText, err := aes.Encrypt(key, text)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("encrypt:%s", cypherText)
	originText, err := aes.Decrypt(key, cypherText)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("decrypt:%s", originText)

}

func TestDecrypt(t *testing.T) {
	ks := "[-9, -126, 104, -32, -65, 80, -2, 39, -23, 48, 63, 20, -90, 27, 44, -64]"
	if strings.Index(ks, "[") == 0 {
		ksr := ks[1 : len(ks)-1]
		log.Printf("ksr:%v", ksr)
		sa := strings.Split(ksr, ",")
		var kes
		for i := 0; i < len(sa); i++ {
			
		}

		for _, v := range sa {
			i, error := strconv.Atoi(strings.TrimSpace(v))
			if error != nil {
				log.Fatal(error)
			}
			log.Printf("%d", i)
		}
	}

	// key := [...]byte{'-9', '-126', '104', '-32', '-65', '80', '-2', '39', '-23', '48', '63', '20', '-90', '27', '44', '-64'}
	// text := "CFA3F485507F35F53EF4CF08FEF3BDB2E49D958DD91A61672E6DD92C5948FDD3"
	// aes := AesHandler{}
	// originText, err := aes.Decrypt(key, text)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("decrypt:%s", originText)

}
