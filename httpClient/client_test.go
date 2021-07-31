package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGetToken(t *testing.T) {
	token := getToken("http://localhost:8003/oauth/token", "client1", "client1-secret")
	fmt.Println(token)
}

func TestGetContent(t *testing.T) {
	content := getContent("http://localhost:8004/client1/api", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiLCJ3ZWJjbGllbnQiXSwiZXhwIjoxNjI3NzI0MjM4LCJqdGkiOiJjNGRlNTA4Zi1mMTdlLTQzOWEtYjJkOS1kYTgyZjI1OTMzM2YiLCJjbGllbnRfaWQiOiJjbGllbnQxIn0.hnrXNkGFdG94tc8s2CftrK2wwvxfR9UskR7pBqTEHsM")
	fmt.Println(content)
}

func TestUpload(t *testing.T) {
	str, err := upload("http://localhost:8005/client2/upload/api", "/home/yoloz/test.txt", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiLCJ3ZWJjbGllbnQiXSwiZXhwIjoxNjI3NzI0MjM4LCJqdGkiOiJjNGRlNTA4Zi1mMTdlLTQzOWEtYjJkOS1kYTgyZjI1OTMzM2YiLCJjbGllbnRfaWQiOiJjbGllbnQxIn0.hnrXNkGFdG94tc8s2CftrK2wwvxfR9UskR7pBqTEHsM")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func TestDownload(t *testing.T) {
	str := download("http://localhost:8005/client2/download", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiLCJ3ZWJjbGllbnQiXSwiZXhwIjoxNjI3NzI0MjM4LCJqdGkiOiJjNGRlNTA4Zi1mMTdlLTQzOWEtYjJkOS1kYTgyZjI1OTMzM2YiLCJjbGllbnRfaWQiOiJjbGllbnQxIn0.hnrXNkGFdG94tc8s2CftrK2wwvxfR9UskR7pBqTEHsM", ".", "test.txt")
	fmt.Println(str)
}

func TestDelete(t *testing.T) {
	str := delete("http://localhost:8005/client2/delete/api", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiLCJ3ZWJjbGllbnQiXSwiZXhwIjoxNjI3NzI0MjM4LCJqdGkiOiJjNGRlNTA4Zi1mMTdlLTQzOWEtYjJkOS1kYTgyZjI1OTMzM2YiLCJjbGllbnRfaWQiOiJjbGllbnQxIn0.hnrXNkGFdG94tc8s2CftrK2wwvxfR9UskR7pBqTEHsM", "test.txt")
	fmt.Println(str)
}
