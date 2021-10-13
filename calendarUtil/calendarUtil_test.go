package calendarUtil

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	fmt.Println(IsLeapYear(2020))
}

func TestGetMonthDay(t *testing.T) {
	i, _ := GetMonthDay(2021, 9)
	fmt.Println(i == 30)
}

func TestGetWeekday(t *testing.T) {
	s, err := GetWeekday(2021, 10, 13)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Compare("星期三", s) == 0)
}
