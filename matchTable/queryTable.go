package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	util "github.com/yoloz/gos/utils/parallelutil"
)

type Config struct {
	CsvDir        string
	SrcId         int
	UserName      string
	Password      string
	Host          string
	Port          int
	DbName        string
	ConnectionMax int
	IdleMax       int
}

//注意方法名大写，就是public
func connect(conf Config) (*sql.DB, error) {
	var db *sql.DB
	var err error
	url := strings.Join([]string{conf.UserName, ":", conf.Password,
		"@tcp(", conf.Host, ":", strconv.Itoa(conf.Port), ")/", conf.DbName, "?charset=utf8"}, "")
	db, err = sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	//设置数据库最大连接数
	db.SetMaxOpenConns(conf.ConnectionMax)
	//设置数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	msg := strings.Join([]string{conf.UserName, "/", conf.Password,
		",", conf.Host, ":", strconv.Itoa(conf.Port), ",", conf.DbName}, "")
	//验证连接
	if err = db.Ping(); err != nil {
		log.Printf("open database %s fail\n", msg)
		return nil, err
	}
	log.Printf("connnect %s success\n", msg)
	return db, nil
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

	db, err := connect(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	os.Mkdir("out", 0774)
	log.Println("create out directory[out]")

	fsys := os.DirFS(conf.CsvDir)         // 以dir为根目录的文件系统，也就是说，后续所有的文件在这目录下
	entries, err := fs.ReadDir(fsys, ".") // 读当前目录
	if err != nil {
		log.Fatal(err)
	}
	size := len(entries)
	worker_num := runtime.NumCPU()
	if size < worker_num {
		worker_num = size
	}
	wp := util.NewPool(worker_num, size).Start()
	wg := sync.WaitGroup{}
	st := time.Now().Unix()
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".csv" && !e.IsDir() {
			f, err := fsys.Open(e.Name()) // 文件名fullname `{dir}/{e.Name()}``
			if err != nil {
				log.Fatal(err)
			}
			wg.Add(1)
			wp.PushTaskFunc(func(args ...interface{}) {
				defer wg.Done()
				compare(args[0].(fs.File), args[1].(int), args[2].(*sql.DB))
			}, f, conf.SrcId, db)
		}
	}
	wg.Wait()
	et := time.Now().Unix()
	log.Printf("all finish cost: %d \n", et-st)
}

func compare(f fs.File, src_id int, db *sql.DB) {
	st := time.Now().Unix()
	csv_stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start operate %v\n", csv_stat.Name())

	outfile_path := filepath.Join(".", "out", csv_stat.Name())
	outfile, err := os.OpenFile(outfile_path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("create out file: %s\n", outfile_path)

	reader := csv.NewReader(f)
	var count int
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if count == 0 {
			count++
			continue
		}
		var database string
		var schema string
		var tableName string
		if len(line) == 3 {
			database = line[0]
			schema = line[1]
			tableName = line[2]
		} else {
			schema = line[0]
			tableName = line[1]
		}
		if !checkExist(src_id, database, schema, tableName, db) {
			_, err := outfile.WriteString(strings.Join(line, ",") + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		count++
	}
	et := time.Now().Unix()
	log.Printf("finish file: %s,line_num: %d,cost: %d\n", csv_stat.Name(), count, et-st)
	f.Close()
	outfile.Close()
}

func checkExist(src_id int, database, schemaName, tableName string, db *sql.DB) bool {
	sql := "SELECT id FROM metaTableDetail mtd WHERE src_id = " + strconv.Itoa(src_id)
	if len(database) > 0 {
		sql += " AND database_name='" + database + "' "
	}
	sql += " AND schema_name ='" + schemaName + "' AND table_name ='" + tableName + "'"
	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("Error:[%s],%v\n", strconv.Itoa(src_id)+","+schemaName+","+tableName, err)
		return false
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		rows.Scan(&id)
		return id != 0
	}
	return false
}
