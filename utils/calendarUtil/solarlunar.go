//reference https://studygolang.com/articles/14740
//reference https://github.com/nosixtools/solarlunar

package calendarUtil

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	MIN_YEAR = 1900
	MAX_YEAR = 2100

	DATELAYOUT   = "2006-01-02"
	STARTDATESTR = "1900-01-30"
)

// LUNAR_INFO 阴历年份的数据,第一行1900-1909,以此类推，倒数二行2090-2099
var LUNAR_INFO = []int{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2,
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977,
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970,
	0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950,
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557,
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0,
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0,
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6,
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570,
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0,
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5,
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930,
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530,
	0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45,
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0,
	0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0,
	0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4,
	0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0,
	0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160,
	0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252,
	0x0d520}

// LunarToSolar 阴历转阳历
func LunarToSolar(year, month, day int, leapMonthFlag bool) string {
	day, offset := dealWithSpecialFebruaryDate(year, month, day)
	lunarTime, err := ParseTime(year, month, day)
	if err != nil {
		fmt.Println(err.Error())
	}
	lunarYear := lunarTime.Year()
	lunarMonth := int(lunarTime.Month())
	lunarDay := lunarTime.Day()
	err = checkLunarDate(lunarYear, lunarMonth, lunarDay, leapMonthFlag)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	for i := MIN_YEAR; i < lunarYear; i++ {
		yearDaysCount := getLunarYearDays(i) // 求阴历某年天数
		offset += yearDaysCount
	}
	//计算该年闰几月
	leapMonth := getLunarLeapMonth(lunarYear)
	if leapMonthFlag && leapMonth != lunarMonth {
		panic("您输入的闰月标志有误！")
	}
	if leapMonth == 0 || (lunarMonth < leapMonth) || (lunarMonth == leapMonth && !leapMonthFlag) {
		for i := 1; i < lunarMonth; i++ {
			tempMonthDaysCount := getLunarMonthDays(lunarYear, uint(i))
			offset += tempMonthDaysCount
		}

		// 检查日期是否大于最大天
		if lunarDay > getLunarMonthDays(lunarYear, uint(lunarMonth)) {
			panic("不合法的农历日期！")
		}
		offset += lunarDay // 加上当月的天数
	} else { //当年有闰月，且月份晚于或等于闰月
		for i := 1; i < lunarMonth; i++ {
			tempMonthDaysCount := getLunarMonthDays(lunarYear, uint(i))
			offset += tempMonthDaysCount
		}
		if lunarMonth > leapMonth {
			temp := getLunarLeapMonthDays(lunarYear) // 计算闰月天数
			offset += temp                           // 加上闰月天数

			if lunarDay > getLunarMonthDays(lunarYear, uint(lunarMonth)) {
				panic("不合法的农历日期！")
			}
			offset += lunarDay
		} else { // 如果需要计算的是闰月，则应首先加上与闰月对应的普通月的天数
			// 计算月为闰月
			temp := getLunarMonthDays(lunarYear, uint(lunarMonth)) // 计算非闰月天数
			offset += temp

			if lunarDay > getLunarLeapMonthDays(lunarYear) {
				panic("不合法的农历日期！")
			}
			offset += lunarDay
		}
	}

	myDate, err := time.Parse(DATELAYOUT, STARTDATESTR)
	if err != nil {
		fmt.Println(err.Error())
	}

	myDate = myDate.AddDate(0, 0, offset)
	return myDate.Format(DATELAYOUT)
}

// dealWithSpecialFebruaryDate 2月日期处理,返回处理后的天数
func dealWithSpecialFebruaryDate(year, month, day int) (int, int) {
	if month == 2 {
		if IsLeapYear(year) {
			if day == 30 {
				return 29, 1
			}
		} else {
			if day == 30 {
				return 28, 2
			}
			if day == 29 {
				return 28, 1
			}
		}
	}
	return day, 0
}

// SolarToLuanr 阳历转阴历.
// bool 月份是否闰月
func SolarToLuanr(year, month, day int) (string, bool) {
	lunarYear, lunarMonth, lunarDay, leapMonth, leapMonthFlag := calculateLunar(year, month, day)
	result := strconv.Itoa(lunarYear) + "-"
	if lunarMonth < 10 {
		result += "0" + strconv.Itoa(lunarMonth) + "-"
	} else {
		result += strconv.Itoa(lunarMonth) + "-"
	}
	if lunarDay < 10 {
		result += "0" + strconv.Itoa(lunarDay)
	} else {
		result += strconv.Itoa(lunarDay)
	}

	if leapMonthFlag && (lunarMonth == leapMonth) {
		return result, true
	} else {
		return result, false
	}
}

//  calculateLunar 计算当前日期的阴历
func calculateLunar(year, month, day int) (lunarYear, lunarMonth, lunarDay, leapMonth int, leapMonthFlag bool) {
	i := 0
	temp := 0
	leapMonthFlag = false
	isLeapYear := false

	myDate, err := ParseTime(year, month, day)
	if err != nil {
		fmt.Println(err.Error())
	}
	startDate, err := time.Parse(DATELAYOUT, STARTDATESTR)
	if err != nil {
		fmt.Println(err.Error())
	}

	offset := daysPay(myDate, startDate)
	for i = MIN_YEAR; i < MAX_YEAR; i++ {
		temp = getLunarYearDays(i) //求当年农历年天数
		if offset-temp < 1 {
			break
		} else {
			offset -= temp
		}
	}
	lunarYear = i

	leapMonth = getLunarLeapMonth(lunarYear) //计算该年闰哪个月

	//设定当年是否有闰月
	if leapMonth > 0 {
		isLeapYear = true
	} else {
		isLeapYear = false
	}

	for i = 1; i <= 12; i++ {
		if i == leapMonth+1 && isLeapYear {
			temp = getLunarLeapMonthDays(lunarYear)
			isLeapYear = false
			leapMonthFlag = true
			i--
		} else {
			temp = getLunarMonthDays(lunarYear, uint(i))
		}
		offset -= temp
		if offset <= 0 {
			break
		}
	}
	offset += temp
	lunarMonth = i
	lunarDay = offset
	return
}

// checkLunarDate 检查阴历是否合法
func checkLunarDate(lunarYear, lunarMonth, lunarDay int, leapMonthFlag bool) error {
	if (lunarYear < MIN_YEAR) || (lunarYear > MAX_YEAR) {
		return errors.New("非法农历年份！")
	}
	if (lunarMonth < 1) || (lunarMonth > 12) {
		return errors.New("非法农历月份！")
	}
	if (lunarDay < 1) || (lunarDay > 30) { // 中国的月最多30天
		return errors.New("非法农历天数！")
	}

	leap := getLunarLeapMonth(lunarYear) // 计算该年应该闰哪个月
	if leapMonthFlag && (lunarMonth != leap) {
		return errors.New("非法闰月！")
	}
	return nil
}

// getLunarMonthDays 计算该月总天数
func getLunarMonthDays(lunarYeay int, month uint) int {
	if month > 31 {
		fmt.Println("error month")
	}
	// 0X0FFFF[0000 {1111 1111 1111} 1111]中间12位代表12个月，1为大月，0为小月
	bit := 1 << (16 - month)
	if ((LUNAR_INFO[lunarYeay-1900] & 0x0FFFF) & bit) == 0 {
		return 29
	} else {
		return 30
	}
}

// getLunarYearDays 计算阴历年的总天数
func getLunarYearDays(year int) int {
	sum := 29 * 12
	for i := 0x8000; i >= 0x8; i >>= 1 {
		if (LUNAR_INFO[year-1900] & 0xfff0 & i) != 0 {
			sum++
		}
	}
	return sum + getLunarLeapMonthDays(year)
}

// getLunarLeapMonthDays	计算阴历年闰月多少天
func getLunarLeapMonthDays(year int) int {
	if getLunarLeapMonth(year) != 0 {
		if (LUNAR_INFO[year-1900] & 0xf0000) == 0 {
			return 29
		} else {
			return 30
		}
	} else {
		return 0
	}
}

// getLunarLeapMonth 计算阴历年闰哪个月 1-12 , 没闰传回 0
func getLunarLeapMonth(year int) int {
	return (int)(LUNAR_INFO[year-1900] & 0xf)
}

// daysPay 当前日期与参考时间的日期补偿
func daysPay(myDate time.Time, startDate time.Time) int {
	subValue := float64(myDate.Unix()-startDate.Unix())/86400.0 + 0.5
	return int(subValue)
}
