package main

import (
	"fmt"
	"time"
)

//判断是否为闰年
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

//获取哪年哪月总天数
func GetMonthDay(year int, month int) (int, error) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, nil
	case 2:
		isLeapYear := IsLeapYear(year)
		if isLeapYear {
			return 29, nil
		} else {
			return 28, nil
		}
	case 4, 6, 9, 11:
		return 30, nil
	default:
		return -1, fmt.Errorf("input month %d is error", month)
	}
}

//获取星期几数据
func GetWeekday(year int, month int, day int) (string, error) {
	t, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%d-%d", year, month, day))
	if err != nil {
		return "", err
	}
	week := t.Weekday()
	switch week {
	case 0:
		return "星期日", nil
	case 1:
		return "星期一", nil
	case 2:
		return "星期二", nil
	case 3:
		return "星期三", nil
	case 4:
		return "星期四", nil
	case 5:
		return "星期五", nil
	default: //6
		return "星期六", nil
	}
}
