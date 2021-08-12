package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	handler "dbHandler/handler"
	"dbHandler/util"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"

	// _ "github.com/ibmdb/go_ibm_db"
	_ "github.com/jackc/pgx/v4/stdlib"
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

type Config struct {
	DbType   DBType
	UserName string
	Password string
	Dsn      string //odbc data sour name
	Host     string
	Port     int
	DbName   string
	PerBatch int
	TotalNum int
}

//注意方法名大写，就是public
func connect(conf Config) (*sql.DB, handler.Handler, error) {
	var db *sql.DB
	var err error
	var handlerImpl handler.Handler
	switch conf.DbType {
	case Mysql:
		handlerImpl = handler.MysqlHandler{}
		url := strings.Join([]string{conf.UserName, ":", conf.Password,
			"@tcp(", conf.Host, ":", strconv.Itoa(conf.Port), ")/", conf.DbName, "?charset=utf8"}, "")
		db, err = sql.Open("mysql", url)
	case Oracle:
		handlerImpl = handler.OracleHandler{}
		url := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%d/%s"`, conf.UserName, conf.Password,
			conf.Host, conf.Port, conf.DbName)
		db, err = sql.Open("godror", url)
	case Mssql:
		handlerImpl = handler.MssqlHandler{}
		url := fmt.Sprintf(`sqlserver://%s:%s@%s:%d?database=%s`, conf.UserName, conf.Password, conf.Host,
			conf.Port, conf.DbName)
		db, err = sql.Open("sqlserver", url)
	case Pgsql:
		handlerImpl = handler.PgsqlHandler{}
		url := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s?sslmode=disable`, conf.UserName, conf.Password, conf.Host,
			conf.Port, conf.DbName)
		db, err = sql.Open("pgx", url)
	case Db2:
		handlerImpl = handler.Db2Handler{}
		url := fmt.Sprintf(`DSN=%s;uid=%s;pwd=%s`, conf.Dsn, conf.UserName, conf.Password)
		db, err = sql.Open("odbc", url)
	}
	if err != nil {
		return nil, nil, err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	//验证连接
	if err = db.Ping(); err != nil {
		fmt.Printf("open database %v fail\n", conf)
		return nil, nil, err
	}
	fmt.Println("connnect success")
	return db, handlerImpl, nil
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()
	var conf Config
	json.NewDecoder(file).Decode(&conf)
	file.Close()
	db, hl, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, e := hl.CreateTable(db)
	if e != nil {
		log.Fatal(e)
	}
	var coroutines = conf.TotalNum / conf.PerBatch
	var startNum = 110101199003071639

	wp := util.NewPool(runtime.NumCPU(), coroutines).Start()
	wg := sync.WaitGroup{}
	st := time.Now().Unix()
	for i := 0; i < coroutines; i++ {
		var start = startNum
		if i != 0 {
			start = startNum + (i * conf.PerBatch) + 1
		}
		wg.Add(1)
		wp.PushTaskFunc(func(args ...interface{}) {
			defer wg.Done()
			hl.BatchInsertSql(args[0].(*sql.DB), args[1].(int), args[2].(int))
		}, db, start, conf.PerBatch)
	}
	wg.Wait()
	et := time.Now().Unix()
	fmt.Printf("all finish %d cost: %d \n", conf.TotalNum, et-st)
}
