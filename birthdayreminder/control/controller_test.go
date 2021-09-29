package control

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	resp, err := http.PostForm("http://localhost:8004/add",
		url.Values{"name": {"张三"}, "timeType": {"1"}, "timeText": {"12-1"}, "sendEmail": {"test@abc.com"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func TestQuery(t *testing.T) {
	resp, err := http.Get("http://localhost:8004/query")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func TestUpdate(t *testing.T) {
	resp, err := http.PostForm("http://localhost:8004/update",
		url.Values{"id": {"1"}, "name": {"里斯"}, "timeType": {"1"}, "timeText": {"12-1"}, "sendEmail": {"test@abc.com"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func TestDel(t *testing.T) {
	resp, err := http.Post("http://localhost:8004/delete",
		"application/x-www-form-urlencoded",
		strings.NewReader("id=1"))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
