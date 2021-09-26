// AES,CBC,PKCS7Padding

package handler

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

type AesHandler struct{}

// aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func (handler AesHandler) Encrypt(key string, text string) (cypher string, err error) {
	pk := []byte(key)
	block, err := aes.NewCipher(pk)
	if err != nil {
		return "", err
	}
	ciphertext := []byte(text)
	blockSize := block.BlockSize()
	originData := pkcs7pad(ciphertext, blockSize)
	iv := []byte(key)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return hex.EncodeToString(crypted), nil
}

func pkcs7pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (handler AesHandler) Decrypt(key string, cypher string) (text string, err error) {
	decode_data, err := hex.DecodeString(cypher)
	if err != nil {
		return "", err
	}
	pk := []byte(key)
	block, err := aes.NewCipher(pk)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	if len(decode_data) < blockSize {
		return "", errors.New("ciphertext too short")

	}
	iv := []byte(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	return string(pkcs7unpad(origin_data)), nil
}

func pkcs7unpad(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
