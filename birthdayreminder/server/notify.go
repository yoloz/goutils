package server

import (
	"log"
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
	month, err := strconv.Atoi(ts[0])
	if err != nil {
		log.Fatalf("month[%s] invalid %v\n", ts[0], err)
	}
	day, err := strconv.Atoi(ts[1])
	if err != nil {
		log.Fatalf("day[%s] invalid %v\n", ts[1], err)
	}
	if birthday.TimeType == 0 {
		solarDate := calendarUtil.LunarToSolar(now.Year(), month, day, false)
		birthday.TimeText = solarDate[4:]
		if solarDay, err = time.ParseInLocation("2006-01-02", solarDate, location); err != nil {
			log.Fatalf("solarDay[%s] invalid %v\n", solarDay, err)
		}
	} else {
		if solarDay, err = time.ParseInLocation("2006-01-02", strconv.Itoa(now.Year())+"-"+birthday.TimeText, location); err != nil {
			log.Fatalf("solarDay[%s] invalid %v\n", solarDay, err)
		}
	}
	duration := solarDay.Sub(now)
	// fmt.Printf("now: %v, birthDay: %v\n", now.Format("2006-01-02 15:04:05"), solarDay.Format("2006-01-02 15:04:05"))
	//2天内即通知
	if 0 < duration && duration < time.Hour*48 {
		SenMail(birthday)
	}
}
