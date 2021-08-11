package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	// mysql "github.com/yoloz/gos/tree/master/dbHandler/dialect"
)

func TestMysql(t *testing.T) {
	dbMsg := DbMsg{0, "test", "dcap123", "192.168.1.116", 3306, "test"}
	db, err := connect(dbMsg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func TestDbMsg(t *testing.T) {
	str := "{\"dbType\":0,\"userName\":\"test\",\"password\":\"dcap123\",\"host\":\"192.168.1.116\",\"port\":3306,\"dbName\":\"test\"}"
	var dbMsg DbMsg
	if err := json.Unmarshal([]byte(str), &dbMsg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", dbMsg)
}
