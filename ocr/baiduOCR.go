package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var maxLen int = 5 * 1024 * 1024

var API_KEY string = ""

var SECRET_KEY string = ""

var OCR_URL string = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic?access_token="

var TOKEN_URL string = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials"

// 通用文字识别（标准版） https://ai.baidu.com/ai-doc/OCR/zk3h7xz52
func GeneralBasic(name string, access_token string) {

	if access_token == "" {
		log.Fatal("access_token is empty")
	}

	fmt.Println("OCR File[" + name + "]")

	file, _ := os.Open(name)
	defer file.Close()

	time.Sleep(2 * time.Second) //qps limit 1

	OCR_URL += access_token
	res, err := http.Post(OCR_URL, "application/x-www-form-urlencoded", strings.NewReader("image="+url.QueryEscape(base64e(file))))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	mapdata := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&mapdata)
	if mapdata["words_result_num"] == nil {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("%s identification failed:%+v\n", file.Name(), data)
		return
	}
	if int(mapdata["words_result_num"].(float64)) == 0 {
		fmt.Println(file.Name() + " identification content is empty...")
		return
	}

	f, err := os.OpenFile(file.Name()+".txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	words_result := mapdata["words_result"].([]interface{})

	for _, obj := range words_result {
		m := obj.(map[string]interface{})
		fmt.Fprintln(w, m["words"])
	}
	w.Flush()
}

// 鉴权认证机制 http://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
// post方式报错:unsupported_grant_type
func Fetch_token() string {
	res, err := http.Get(TOKEN_URL + "&client_id=" + API_KEY + "&client_secret=" + SECRET_KEY)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	mapdata := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&mapdata)
	// json.Unmarshal(data, &mapdata)
	for key, value := range mapdata {
		// fmt.Println("key:", key, " => value :", value)
		if key == "access_token" {
			return value.(string)
		}
		if key == "error_description" {
			log.Fatal(value)
		}
	}
	return ""
}

// base64 encode
func base64e(file *os.File) string {
	sourcebuffer := make([]byte, maxLen)
	n, _ := file.Read(sourcebuffer)
	return base64.StdEncoding.EncodeToString(sourcebuffer[:n])
}

// base64 decode
func base64d(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
