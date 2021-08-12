package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestDbMsg(t *testing.T) {
	str := "{\"dbType\":0,\"userName\":\"test\",\"password\":\"test\",\"dsn\":\"\",\"host\":\"127.0.0.1\",\"port\":3306,\"dbName\":\"test\",\"perBatch\":10,\"totalNum\":100}"
	var conf Config
	if err := json.Unmarshal([]byte(str), &conf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", conf)
}

func TestMysql(t *testing.T) {
	conf := Config{0, "test", "test", "", "127.0.0.1", 3306, "test", 100, 100}
	db, ml, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := ml.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	result, err := ml.BatchInsertSql(db, 110101199003071639, 1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}

func TestOracle(t *testing.T) {
	conf := Config{1, "test", "test", "", "127.0.0.1", 1521, "test", 100, 100}
	db, ol, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := ol.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	result, err := ol.BatchInsertSql(db, 110101199003071639, 200)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}

func TestMssql(t *testing.T) {
	conf := Config{3, "test", "test", "", "127.0.0.1", 1433, "test", 100, 100}
	db, msl, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := msl.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	result, err := msl.BatchInsertSql(db, 110101199003071639, 500)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}

func TestPgsql(t *testing.T) {
	conf := Config{2, "test", "test", "", "127.0.0.1", 5432, "test", 100, 100}
	db, pgl, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := pgl.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	result, err := pgl.BatchInsertSql(db, 110101199003071639, 2000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}

func TestDb2(t *testing.T) {
	conf := Config{4, "test", "test", "133DB2", "", 0, "", 100, 100}
	db, dl, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := dl.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	result, err := dl.BatchInsertSql(db, 110101199003071639, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())
}
