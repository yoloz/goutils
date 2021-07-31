package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	jc := make(map[string]string)
	json.NewDecoder(file).Decode(&jc)

	access_token_uri := jc["access_token_uri"]
	client_id := jc["client_id"]
	client_secret := jc["client_secret"]
	repeat_count, err := strconv.Atoi(jc["repeat_count"])
	if err != nil {
		log.Fatal(err)
	}
	token := getToken(access_token_uri, client_id, client_secret)
	authoritization := "bearer " + token
	client1_api_uri := jc["client1_api_uri"]
	upload_uri := jc["upload_uri"]
	upload_path := jc["upload_path"]
	download_uri := jc["download_uri"]
	download_path := jc["download_path"]
	delete_uri := jc["delete_uri"]
	for i := 0; i < repeat_count; i++ {
		client1_resp := getContent(client1_api_uri, authoritization)
		fmt.Println("client1 resp:\r\n" + client1_resp)
		upload_resp, err := upload(upload_uri, upload_path, authoritization)
		if err != nil {
			fmt.Printf("updaload %s fail\r\n", upload_path)
		} else {
			fmt.Println("upload resp:\r\n" + upload_resp)
			download_resp := download(download_uri, authoritization, download_path, upload_resp)
			fmt.Println("download resp:\r\n" + download_resp)
		}
		if len(delete_uri) > 0 {
			delete_resp := delete(delete_uri, authoritization, upload_resp)
			fmt.Println("delete resp:\r\n" + delete_resp)
		}
		time.Sleep(time.Second)
	}

}

/**
上传文件：

Content-Type:multipart/form-data; boundary=----WebKitFormBoundaryExT8avmSnrECoDbP

------WebKitFormBoundaryExT8avmSnrECoDbP
Content-Disposition: form-data; name="name"

qwe
------WebKitFormBoundaryExT8avmSnrECoDbP
Content-Disposition: form-data; name="pwd"

123
------WebKitFormBoundaryExT8avmSnrECoDbP
Content-Disposition: form-data; name="icon"; filename="0fbc751ff63fa8cf3302b03889b9421e65d6592301"
Content-Type: application/octet-stream


------WebKitFormBoundaryExT8avmSnrECoDbP--
**/

func delete(uri string, authorization string, filename string) string {

	client := http.DefaultClient
	req, err := http.NewRequest("POST", uri, strings.NewReader("filename="+filename))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(resp_body)
}

func upload(uri string, file string, authorization string) (string, error) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作，名称必须file，否则报错
	//org.springframework.web.multipart.support.MissingServletRequestPartException: Required request part 'file' is not present
	fileWriter, err := bodyWriter.CreateFormFile("file", path.Base(file))
	if err != nil {
		fmt.Println("error writing to buffer")
		return "", err
	}

	//打开文件句柄操作
	fh, err := os.Open(file)
	if err != nil {
		fmt.Println("error opening file")
		return "", err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}

	contentType := bodyWriter.FormDataContentType()
	// 这里需要关闭bodyWriter,不可defer方式，否则报错如下
	//org.apache.tomcat.util.http.fileupload.MultipartStream$MalformedStreamException: Stream ended unexpectedly
	bodyWriter.Close()

	// resp, err := http.Post(uri, contentType, bodyBuf)
	client := http.DefaultClient
	req, err := http.NewRequest("POST", uri, bodyBuf)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resp_body), nil
}

func download(uri string, authorization string, dwdir string, filename string) string {

	client := http.DefaultClient
	req, err := http.NewRequest("POST", uri, strings.NewReader("filename="+filename))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authorization)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	dwfile := path.Join(dwdir, filename)
	out, err := os.Create(dwfile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return dwfile
}

// curl -H "Authorization: Bearer ACCESS_TOKEN" http://192.168.90.124:8004/client1/api
func getContent(uri string, authorization string) string {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", authorization)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return string(content)
}

// http://192.168.90.124:8003/oauth/token?grant_type=client_credentials&client_id=client1&client_secret=client1-secret
// {
//     "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiLCJ3ZWJjbGllbnQiXSwiZXhwIjoxNjI3NzE0MzEyLCJqdGkiOiI3MDQ1MDUxYS1mZjZlLTQzY2YtYTMyMC1hN2RiZWE3ZGM2OTkiLCJjbGllbnRfaWQiOiJjbGllbnQxIn0.Vd7IKmDkM0EgATaobsO7x_lw_tLYn9zt_9HYfiEwnEk",
//     "token_type": "bearer",
//     "expires_in": 3599,
//     "scope": "read write webclient",
//     "jti": "7045051a-ff6e-43cf-a320-a7dbea7dc699"
// }
func getToken(access_token string, clientId string, clientSecret string) string {
	res, err := http.Get(access_token + "?grant_type=client_credentials" + "&client_id=" + clientId + "&client_secret=" + clientSecret)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	mapdata := make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&mapdata)
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
