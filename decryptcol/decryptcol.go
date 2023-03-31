package main

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wxnacy/wgo/arrays"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"

	"github.com/yoloz/gos/utils/cryptoutil"
)

// Db数据库连接池
var DB *sql.DB

type DBType uint8

const (
	Mysql DBType = iota
	Oracle
	Pgsql
	Mssql
)

type Config struct {
	DbType    DBType
	Host      string
	Port      int
	UserName  string
	Password  string
	DbName    string
	TableName string
	Columns   []string //加密列
	Indexes   []string //更新语句的索引列
	QWhere    string   //查询条件
	UWhere    string   //更新条件,和索引列对应 "c1='%s' and c2=%s"
	EncKey    string   //加密密钥
	HexCode   bool     //值是否hex处理
}

// 注意方法名大写，就是public
func connect(conf Config) (*sql.DB, error) {
	var db *sql.DB
	var err error
	switch conf.DbType {
	case Mysql:
		url := strings.Join([]string{conf.UserName, ":", conf.Password,
			"@tcp(", conf.Host, ":", strconv.Itoa(conf.Port), ")/", conf.DbName, "?charset=utf8"}, "")
		db, err = sql.Open("mysql", url)
	case Oracle:
		url := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%d/%s"`, conf.UserName, conf.Password,
			conf.Host, conf.Port, conf.DbName)
		db, err = sql.Open("godror", url)
	case Mssql:
		url := fmt.Sprintf(`sqlserver://%s:%s@%s:%d?database=%s`, conf.UserName, conf.Password, conf.Host,
			conf.Port, conf.DbName)
		db, err = sql.Open("sqlserver", url)
	case Pgsql:
		url := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s?sslmode=disable`, conf.UserName, conf.Password, conf.Host,
			conf.Port, conf.DbName)
		db, err = sql.Open("pgx", url)
	}
	if err != nil {
		return nil, err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置数据库最大闲置连接数
	db.SetMaxIdleConns(2)
	//验证连接
	if err = db.Ping(); err != nil {
		log.Printf("open database %s:%d fail\n", conf.Host, conf.Port)
		return nil, err
	}
	log.Println("connnect success")
	return db, nil
}

// return query and update sql
func sqlfactory(conf Config) (string, string) {
	var qstr, ustr string
	switch conf.DbType {
	case Mysql:
		qstr = "select "
		for i := 0; i < len(conf.Columns); i++ {
			qstr += conf.Columns[i]
			if i != len(conf.Columns)-1 {
				qstr += ","
			}
		}
		for i := 0; i < len(conf.Indexes); i++ {
			qstr += ("," + conf.Indexes[i])
		}
		qstr += (" from " + conf.TableName)
		if conf.QWhere != "" {
			qstr += (" where " + conf.QWhere)
		}
		log.Printf("query sql:%s\n", qstr)

		ustr = ("update " + conf.TableName + " set ")
		for i := 0; i < len(conf.Columns); i++ {
			ustr += (conf.Columns[i] + "='%s'")
			if i != len(conf.Columns)-1 {
				ustr += ","
			}
		}
		ustr += (" where " + conf.UWhere)
	case Oracle:
	case Mssql:
	case Pgsql:
		return "", ""
	}

	return qstr, ustr
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
	qstr, ustr := sqlfactory(conf)

	rows, err := db.Query(qstr)
	if err != nil {
		log.Panicf("Error:[%s],%v\n", qstr, err)
	}
	defer rows.Close()

	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}

	for rows.Next() {
		err := rows.Scan(values...) //开始读行，Scan函数只接受指针变量
		if err != nil {
			panic(err)
		}
		// m := make(map[string]string) //用于存放1行的 [键/值] 对
		//简化处理，查询语句已经指定了列的顺序
		m := make([]interface{}, count)
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			b, _ := raw_value.([]byte)
			v := string(b) //将raw数据转换成字符串
			if arrays.ContainsString(conf.Columns, colName) > -1 {
				var vb []byte
				if conf.HexCode {
					vb, err = hex.DecodeString(v)
					if err != nil {
						panic(err)
					}
				}
				aes := cryptoutil.AES{}
				if strings.Index(conf.EncKey, "[") == 0 {
					ksr := conf.EncKey[1 : len(conf.EncKey)-1]
					sa := strings.Split(ksr, ",")
					var keys = make([]byte, len(sa))

					for i := 0; i < len(sa); i++ {
						n, error := strconv.Atoi(strings.TrimSpace(sa[i]))
						if error != nil {
							log.Fatal(error)
						}
						//上述字节来自java,而golang中byte(0-255),负数+256
						keys[i] = byte(n)
					}
					b, err := aes.DecryptECBImpl(keys, vb)
					if err != nil {
						panic(err)
					}
					v = string(b)
				} else {
					b, err := aes.DecryptECB(conf.EncKey, string(vb))
					if err != nil {
						panic(err)
					}
					v = string(b)
				}
			}
			// m[colName] = v //colName是键，v是值
			m[i] = v
		}
		upsql := fmt.Sprintf(ustr, m...)
		_, uerr := db.Exec(upsql) //增、删、改就靠这一条命令就够了，很简单
		if uerr != nil {
			log.Panicf("update:%s fail %v\n", upsql, uerr)
		}

	}
}
