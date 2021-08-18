package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBase64(t *testing.T) {
	//读原图片
	h, _ := Home()
	ff, _ := os.Open(filepath.Join(h, "test.png"))
	defer ff.Close()
	sourcestring := base64e(ff)

	urle := url.QueryEscape(sourcestring)
	urld, _ := url.QueryUnescape(urle)
	fmt.Println(urle, strings.Compare(sourcestring, urld))
	//写入临时文件
	ioutil.WriteFile(filepath.Join(h, "test.png.txt"), []byte(sourcestring), 0667)
	//读取临时文件
	cc, _ := ioutil.ReadFile(filepath.Join(h, "test.png.txt"))

	dist, _ := base64d(string(cc))
	//写入新文件
	f, _ := os.OpenFile(filepath.Join(h, "test1.png"), os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}

func TestFetch_token(t *testing.T) {
	fmt.Println(Fetch_token())
}

func TestGeneralBasic(t *testing.T) {
	access_token := Fetch_token()
	GeneralBasic("test.jpg", access_token)
}

func TestMapType(t *testing.T) {
	jsonBuf := `
	{
		"words_result":[
		{
		"words":"中午,胡小懒跟公司的同事去附近的餐馆吃饭,附近所有的餐馆"
		},
		{
		"words":"都人满为患,往往需要等位,午餐时间至少需要45分钟。吃完午饭回"
		}
		],
		"log_id":1402434780670197760,
		"words_result_num":2
		}
   `

	//创建一个map
	m := make(map[string]interface{}, 3)
	//第二个参数要地址传递
	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m = %+v\n", m)

	var str string

	//类型断言, 值，它是value类型
	for key, value := range m {
		//fmt.Printf("%v ============> %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string, value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool, value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64, value = %f\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string, value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]interface, value = %v\n", key, data)
		}

	}

}
