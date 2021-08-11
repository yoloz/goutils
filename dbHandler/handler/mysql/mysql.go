package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func CreateTable(db *sql.DB) (sql.Result, error) {
	str := `DROP TABLE IF EXISTS test.person; 
	CREATE TABLE test.person (sfzh varchar(100) NOT NULL,birth DATE NOT NULL,
	age INT NOT NULL,ip varchar(100) NOT NULL,post INT NOT NULL,PRIMARY KEY (sfzh)) 
	ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
	result, err := db.Exec(str)
	return result, err
}

//拼接insert语句
func BatchInsertSql(db *sql.DB, startNum int64, batchNum int) (sql.Result, error) {
	st := time.Now().Unix()
	prefix := "INSERT INTO person VALUES"
	var sfzh string
	var birth string
	var age int
	var ip string
	var post int
	var insertBuf string = prefix

	for i := 1; i <= batchNum; i++ {
		sfzh = strconv.FormatInt(startNum+int64(i), 10)
		birth = MakeRandDate()
		age = int(MakeRandInt(15, 80))
		ip = MakeRandIPV4()
		post = int(MakeRandInt(100000, 999999))
		insertBuf += fmt.Sprintf(" (%s, '%s','%d',%s, %d) ", sfzh, birth, age, ip, post)
	}
	insertBuf += ";"

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
