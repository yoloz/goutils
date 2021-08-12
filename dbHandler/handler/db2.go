package handler

import (
	"database/sql"
	"dbHandler/util"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Db2Handler struct{}

func (handler Db2Handler) CreateTable(db *sql.DB) (sql.Result, error) {
	var result sql.Result
	var err error
	var count int
	row := db.QueryRow("select COUNT(*) from SYSCAT.TABLES where TRIM(TABNAME) = 'PERSON';")
	if er := row.Scan(&count); er != nil {
		log.Fatal(er)
	}
	fmt.Printf("table exist:%d\n", count)
	if count > 0 {
		_, err = db.Exec("DROP TABLE DB2INST1.PERSON;")
		// _, err = db.Exec("TRUNCATE TABLE DB2INST1.PERSON immediate")
		if err != nil {
			return nil, err
		}
		// } else {
		// 	return nil, errors.New("table is not exist")
	}
	str := `CREATE TABLE DB2INST1.PERSON (
		sfzh VARCHAR(50) NOT NULL,
		birth DATE NOT NULL,
		age INTEGER NOT NULL,
		ip VARCHAR(50) NOT NULL,
		post INTEGER NOT NULL,
		PRIMARY KEY (sfzh)
		);`
	result, err = db.Exec(str)
	return result, err
	// return nil, nil
}

//拼接insert语句
func (handler Db2Handler) BatchInsertSql(db *sql.DB, startNum int, batchNum int) (sql.Result, error) {
	st := time.Now().Unix()
	prefix := "INSERT INTO DB2INST1.PERSON VALUES"
	var sfzh string
	var birth string
	var age int
	var ip string
	var post int
	var insertBuf string = prefix

	for i := 1; i <= batchNum; i++ {
		sfzh = strconv.FormatInt(int64(startNum+i), 10)
		birth = util.MakeRandDate()
		age = int(util.MakeRandInt(15, 80))
		ip = util.MakeRandIPV4()
		post = int(util.MakeRandInt(100000, 999999))
		insertBuf += fmt.Sprintf(" ('%s', '%s',%d,'%s', %d) ", sfzh, birth, age, ip, post)
		if i != batchNum {
			insertBuf += ","
		} else {
			insertBuf += ";"
		}
	}
	var rs sql.Result
	var err error

	for i := 5; i > 0; i-- {
		fmt.Println(insertBuf)
		rs, err = db.Exec(insertBuf)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			break
		}
	}
	et := time.Now().Unix()
	fmt.Printf("startNum:[%d]  batchNum:[%d]  cost:[%ds]\n", startNum, batchNum, et-st)
	return rs, err
}
