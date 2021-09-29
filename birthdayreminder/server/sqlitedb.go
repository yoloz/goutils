//sample from https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go

package server

import (
	"bufio"
	"container/list"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Birthday struct {
	// `json:"-"` // 表示不进行序列化
	Id        int    `json:"id"` //primary key
	Name      string `json:"name"`
	TimeType  int    `json:"timeType"` //农历:０,公历:1
	TimeText  string `json:"timeText"` //2-12
	SendEmail string `json:"sendEmail"`
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./reminder.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitDb() {

	os.Remove("./reminder.db")

	db := openDb()
	defer db.Close()

	sqlStmt := `CREATE TABLE birthday(
		id             INTEGER   PRIMARY KEY  AUTOINCREMENT,
		name           TEXT      NOT NULL,
		timeType       INTEGER   NOT NULL,
		timeText       TEXT      NOT NULL,
		sendEmail      TEXT      NOT NULL
	 );
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertDb(birthDay Birthday) error {

	db := openDb()
	defer db.Close()

	_, err := db.Exec("insert into birthday(name,timeType,timeText,sendEmail) values" +
		fmt.Sprintf(" ('%s', %d, '%s', '%s') ", birthDay.Name, birthDay.TimeType, birthDay.TimeText,
			birthDay.SendEmail))
	return err
}

func UpdateDb(birthday Birthday) error {

	db := openDb()
	defer db.Close()

	_, err := db.Exec(fmt.Sprintf(
		"update birthday set name='%s', timeType=%d, timeText='%s', sendEmail='%s' where id=%d",
		birthday.Name, birthday.TimeType, birthday.TimeText, birthday.SendEmail, birthday.Id))
	return err
}

func DeleteDb(id int) error {

	db := openDb()
	defer db.Close()

	_, err := db.Exec(fmt.Sprintf("delete from birthday where id=%d", id))
	return err
}

func QueryAll() *list.List {

	db := openDb()
	defer db.Close()

	rows, err := db.Query("select id,name,timeType,timeText,sendEmail from birthday")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	list := list.New()
	for rows.Next() {
		var id int
		var name string
		var timeType int
		var timeText string
		var sendEmail string
		err = rows.Scan(&id, &name, &timeType, &timeText, &sendEmail)
		if err != nil {
			log.Fatal(err)
		}
		birthday := Birthday{
			Id:        id,
			Name:      name,
			TimeType:  timeType,
			TimeText:  timeText,
			SendEmail: sendEmail,
		}
		list.PushBack(birthday)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func ReadConf() *list.List {
	file, err := os.Open("conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	list := list.New()
	for scanner.Scan() {
		lineText := scanner.Text()
		fileds := strings.Fields(lineText)
		if len(fileds) < 4 {
			log.Println("format error[" + lineText + "]")
		}
		t, err := strconv.Atoi(fileds[1])
		if err != nil {
			log.Fatalln(err)
		}
		list.PushBack(Birthday{0, fileds[0], t, fileds[2], fileds[3]})
	}
	return list
}
