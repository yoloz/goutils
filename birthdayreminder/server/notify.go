package server

import (
	"strconv"
	"strings"
	"time"

	"github.com/yoloz/gos/utils/calendarUtil"
)

func EmailNotify() {
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	for {
		l := QueryAll()
		now := time.Now()
		now.In(timeLocation)
		for e := l.Front(); e != nil; e = e.Next() {
			birthDay, _ := e.Value.(Birthday)
			notifyEmail(birthDay, now, timeLocation)
		}
		time.Sleep(time.Hour * 24)
		// time.Sleep(time.Minute * 10)
	}
}

func notifyEmail(birthday Birthday, now time.Time, location *time.Location) {
	var solarDay time.Time
	ts := strings.Split(birthday.TimeText, "-")
	month, _ := strconv.Atoi(ts[0])
	day, _ := strconv.Atoi(ts[1])
	if birthday.TimeType == 0 {
		solar := calendarUtil.LunarToSolar(now.Year(), month, day, false)
		birthday.TimeText = strconv.Itoa(solar[1]) + "-" + strconv.Itoa(solar[2])
		solarDay = time.Date(solar[0], time.Month(solar[1]), solar[2], 12, 0, 0, 0, location)
	} else {
		solarDay = time.Date(now.Year(), time.Month(month), day, 12, 0, 0, 0, location)
	}
	duration := solarDay.Sub(now)
	// fmt.Printf("now: %v, birthDay: %v\n", now.Format("2006-01-02 15:04:05"), solarDay.Format("2006-01-02 15:04:05"))
	//2天内即通知
	if 0 < duration && duration < time.Hour*48 {
		SenMail(birthday)
	}
}
