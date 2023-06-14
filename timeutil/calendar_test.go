package timeutil

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestIsLeapYear(t *testing.T) {
	fmt.Println(IsLeapYear(2020))
}

func TestGetMonthDay(t *testing.T) {
	i, _ := GetMonthDay(2021, 9)
	fmt.Println(i == 30)
}

func TestGetWeekday(t *testing.T) {
	var (
		week time.Weekday
		err  error
	)
	if week, err = GetWeekday(2021, 10, 1); err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Compare("星期五", Weekday_zh(week)) == 0)
	if week, err = GetWeekday(2021, 10, 13); err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Compare("星期三", Weekday_zh(week)) == 0)
}
