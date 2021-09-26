package main

import (
	"bufio"
	"cryptoHandler/handler"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	var key string
	var mode int
	var text string
	fmt.Println("AES加密==>")
	fmt.Printf("请输入密钥(长度16|24|32): ")
	fmt.Scanln(&key)
	leng := len(key)
	if leng%16 != 0 {
		log.Fatal(errors.New("密钥长度不匹配"))
	}
	fmt.Printf("请选择加密(1),解密(2): ")
	fmt.Scanln(&mode)
	if mode != 1 && mode != 2 {
		log.Fatal("操作：“+mode+”未定义")
	}
	fmt.Printf("请输入待处理字符: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()
		break
	}
	aeshandler := handler.AesHandler{}
	if mode == 1 {
		cipherText, err := aeshandler.Encrypt(key, text)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("加密后的字符: " + cipherText)
		return
	}
	originText, err := aeshandler.Decrypt(key, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("解密后的字符: " + originText)
}
