// AES,CBC,ECB,CFB PKCS7Padding

// 在PKCS5Padding中，明确定义Block的大小是8位，而在PKCS7Padding定义中，对于块的大小是不确定的，可以在1-255之间（块长度超出255的尚待研究），填充值的算法是一样的

// aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.

package cryptoutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type AES struct{}

// =======================================CBC========================================

func (aes_ AES) EncryptCBC(key string, text string) (cypher string, err error) {
	return aes_.EncryptCBCImpl([]byte(key), []byte(key), []byte(text))
}

func (aes_ AES) EncryptCBCImpl(pk []byte, iv []byte, ciphertext []byte) (cypher string, err error) {
	block, err := aes.NewCipher(pk)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	originData := pkcs7pad(ciphertext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return hex.EncodeToString(crypted), nil
}

func (aes_ AES) DecryptCBC(key string, cypher string) (text string, err error) {
	return aes_.DecryptCBCImpl([]byte(key), []byte(key), cypher)
}

func (aes_ AES) DecryptCBCImpl(pk []byte, iv []byte, cypher string) (text string, err error) {
	decode_data, err := hex.DecodeString(cypher)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(pk)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	if len(decode_data) < blockSize {
		return "", errors.New("ciphertext too short")

	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	return string(pkcs7unpad(origin_data)), nil
}

func pkcs7pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7unpad(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// ================================ECB===============================

func (aes_ AES) EncryptECB(key string, text string) (cypher string, err error) {
	bs, err := aes_.EncryptECBImpl([]byte(key), []byte(text))
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (aes_ AES) EncryptECBImpl(key []byte, cyphertext []byte) (encrypted []byte, err error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	length := (len(cyphertext) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, cyphertext)
	pad := byte(len(plain) - len(cyphertext))
	for i := len(cyphertext); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(cyphertext); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

func (aes_ AES) DecryptECB(key string, cypher string) (text string, err error) {
	bs, err := aes_.DecryptECBImpl([]byte(key), []byte(cypher))
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (aes_ AES) DecryptECBImpl(key []byte, encrypted []byte) (decrypted []byte, err error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}

// ==========================================CFB====================================

func (aes_ AES) EncryptCFB(key, text string) (encrypted string, err error) {
	bs, err := aes_.EncryptCFBImpl([]byte(key), []byte(text))
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (aes_ AES) EncryptCFBImpl(key []byte, origData []byte) (encrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, nil
}

func (aes_ AES) DecryptCFB(key string, cypher string) (text string, err error) {
	bs, err := aes_.DecryptCFBImpl([]byte(key), []byte(cypher))
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (aes_ AES) DecryptCFBImpl(key []byte, encrypted []byte) (decrypted []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	decrypted = make([]byte, aes.BlockSize+len(encrypted))
	iv := decrypted[:aes.BlockSize]
	decrypted = decrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, encrypted)
	return encrypted, nil
}
