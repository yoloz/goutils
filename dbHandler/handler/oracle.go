package handler

import (
	"database/sql"
	"dbHandler/util"
	"fmt"
	"strconv"
	"time"
)

type OracleHandler struct{}

func (handler OracleHandler) CreateTable(db *sql.DB) (sql.Result, error) {
	var result sql.Result
	var err error
	_, err = db.Exec(`BEGIN
	EXECUTE IMMEDIATE 'DROP TABLE PERSON';
  EXCEPTION
	WHEN OTHERS THEN
	  NULL;
  END;`)
	if err != nil {
		return nil, err
	}
	str := `CREATE TABLE PERSON (
		SFZH VARCHAR2(100) NOT NULL,
		BIRTH DATE NOT NULL,
		AGE INTEGER NOT NULL,
		IP VARCHAR2(100) NOT NULL,
		POST INTEGER NOT NULL,
		PRIMARY KEY (SFZH)
		)` //last ; will cause missing option
	result, err = db.Exec(str)
	return result, err
}

//拼接insert语句,适合于Oralce 9i以上版本
func (handler OracleHandler) BatchInsertSql(db *sql.DB, startNum int, batchNum int) (sql.Result, error) {
	st := time.Now().Unix()
	prefix := "BEGIN INSERT ALL INTO PERSON VALUES"
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
		insertBuf += fmt.Sprintf(" ('%s', to_date('%s','yyyy-mm-dd'),%d,'%s', %d) ", sfzh, birth, age, ip, post)
		if i != batchNum {
			insertBuf += " INTO PERSON VALUES"
		} else {
			insertBuf += "SELECT 1 FROM DUAL; END;"
		}
	}
	var rs sql.Result
	var err error

	for i := 5; i > 0; i-- {
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
