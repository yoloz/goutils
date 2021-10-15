package calendarUtil

import (
	"fmt"
	"time"
)

// IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

// GetMonthDay 获取哪年哪月总天数
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

// GetWeekday 获取星期几数据
func GetWeekday(year int, month int, day int) (time.Weekday, error) {
	t, err := ParseTime(year, month, day)
	if err != nil {
		return -1, err
	}
	return t.Weekday(), nil
}

// ParseTime 年月日格式化
func ParseTime(year, month, day int) (time.Time, error) {
	df := "2006-"
	if month < 10 {
		df += "1-"
	} else {
		df += "01-"
	}
	if day < 10 {
		df += "2"
	} else {
		df += "02"
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(df, fmt.Sprintf("%d-%d-%d", year, month, day), timeLocation)
}

// Weekday_zh 输出中文星期
func Weekday_zh(week time.Weekday) string {
	switch week {
	case 0:
		return "星期日"
	case 1:
		return "星期一"
	case 2:
		return "星期二"
	case 3:
		return "星期三"
	case 4:
		return "星期四"
	case 5:
		return "星期五"
	default: //6
		return "星期六"
	}
}
