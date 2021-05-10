package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"birthdayreminder/convert"
)

type conf struct {
	Name  string
	Month int
	Day   int
	Type  int //农历:０,公历:1
}

func readConf() *list.List {
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
		month, err := strconv.Atoi(fileds[1])
		if err != nil {
			log.Fatalln(err)
		}
		day, err := strconv.Atoi(fileds[2])
		if err != nil {
			log.Fatalln(err)
		}
		t, err := strconv.Atoi(fileds[3])
		if err != nil {
			log.Fatalln(err)
		}
		list.PushBack(conf{fileds[0], month, day, t})
	}
	return list
}

const (
	// 邮件服务器地址
	smtpMailHost = "smtp.126.com"
	// 端口
	smtpMailPort = "25"
	// 发送邮件用户账号
	smtpMailUser = "xxx@126.com"
	// 授权密码
	smtpMailPwd = "xxx"
	// TO
	smtpMailTo = "xxx@outlook.com"
)

func msg(cf conf) string {
	var build strings.Builder
	build.WriteString(cf.Name)
	build.WriteString("的生日在近期[")
	build.WriteString(strconv.Itoa(cf.Month))
	build.WriteString("-")
	build.WriteString(strconv.Itoa(cf.Day))
	build.WriteString("]")
	return build.String()
}

func senMail(cf conf) {
	auth := smtp.PlainAuth("", smtpMailUser, smtpMailPwd, smtpMailHost)
	to := []string{smtpMailTo}
	msg := []byte("To: " + smtpMailTo + "\r\n" +
		"Subject: Birthday Reminder!\r\n" +
		"\r\n" +
		msg(cf) + "\r\n")
	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpMailHost, smtpMailPort), auth, smtpMailUser, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// l := readConf()
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	for {
		l := readConf()
		now := time.Now()
		now.In(timeLocation)
		for e := l.Front(); e != nil; e = e.Next() {
			var solarDay time.Time
			cf, _ := e.Value.(conf)
			if cf.Type == 0 {
				solar := convert.LunarToSolar(now.Year(), cf.Month, cf.Day, false)
				cf.Month = solar[1]
				cf.Day = solar[2]
				solarDay = time.Date(solar[0], time.Month(solar[1]), solar[2], 12, 0, 0, 0, timeLocation)
			} else {
				solarDay = time.Date(now.Year(), time.Month(cf.Month), cf.Day, 12, 0, 0, 0, timeLocation)
			}
			duration := solarDay.Sub(now)
			// fmt.Printf("now: %v, birthDay: %v\n", now.Format("2006-01-02 15:04:05"), solarDay.Format("2006-01-02 15:04:05"))
			if 0 < duration && duration < time.Hour*2*24 {
				senMail(cf)
				fmt.Printf("time:%v,msg:%v\n", now.Format("2006-01-02 15:04:05"), msg(cf))
			}
		}
		time.Sleep(time.Hour * 24)
		// time.Sleep(time.Minute * 10)
	}
}
