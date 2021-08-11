package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//Db数据库连接池
var DB *sql.DB

type DBType uint8

const (
	Mysql DBType = iota
	Oracle
	Pgsql
	Mssql
	Db2
)

type DbMsg struct {
	DbType   DBType
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
}

//注意方法名大写，就是public
func connect(dbMsg DbMsg) (*sql.DB, error) {
	var db *sql.DB
	var err error
	switch dbMsg.DbType {
	case Mysql:
		url := strings.Join([]string{dbMsg.UserName, ":", dbMsg.Password,
			"@tcp(", dbMsg.Host, ":", strconv.Itoa(dbMsg.Port), ")/", dbMsg.DbName, "?charset=utf8"}, "")
		db, err = sql.Open("mysql", url)
		if err != nil {
			return nil, err
		}
	case Oracle:

	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(20)
	//设置数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	//验证连接
	if err = db.Ping(); err != nil {
		fmt.Printf("open database %v fail\n", dbMsg)
		return nil, err
	}
	fmt.Println("connnect success")
	return db, nil
}

func main() {}
