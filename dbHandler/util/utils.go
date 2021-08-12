package util

import (
	"math/rand"
	"strconv"
	"time"
)

//随机字符串种子
const STRCHAR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&*+-./:;<=>?@[]^_{|}~"

var CITYS = []string{"ChengDu", "KunMing", "XiAn", "LaSa", "JiNan", "NanJing", "HangZhou", "FuZhou", "GuangZhou",
	"HaiKou", "HaErBin", "ChangChun", "ShenYang", "ZhengZhou", "HeFei", "WuHan", "ChongQing", "BeiJing", "ShangHai"}

//随机生成一个整型数据
func MakeRandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

//随机生成一个浮点型数据
func MakeRandFloat(base int64) float32 {
	return rand.Float32() * float32(base)
}

//随机生成一个双精度浮点型数据
func MakeRandDouble(base int64) float64 {
	return rand.Float64() * float64(base)
}

//随机生成一个汉字字符串  入参：字符串长度
func MakeChineseString(length int) string {
	a := make([]rune, length)
	for i := range a {
		a[i] = rune(MakeRandInt(19968, 40869))
	}
	return string(a)
}

//随机生成一个IPV4
func MakeRandIPV4() string {
	i1 := MakeRandInt(10, 254)
	i2 := MakeRandInt(10, 254)
	i3 := MakeRandInt(10, 254)
	i4 := MakeRandInt(10, 254)
	return strconv.FormatInt(i1, 10) + "." + strconv.FormatInt(i2, 10) + "." + strconv.FormatInt(i3, 10) + "." + strconv.FormatInt(i4, 10)
}

//随机生成一个字符串（指定字符种子） 入参：（字符串长度，长度是否随机）
func MakeRandString(length int64, bRegular bool) string {
	var size int64
	if bRegular {
		size = length
	} else {
		size = MakeRandInt(1, length)
	}

	str := make([]byte, size)
	for i := 0; i < int(size); i++ {
		index := MakeRandInt(0, int64(len(STRCHAR)))
		str[i] = STRCHAR[index]
	}
	return string(str)
}

//随机生成一个字符串（任意字符） 入参：字符串长度
func MakeRandString2(length int) string {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		index := MakeRandInt(0, 127)
		str[i] = byte(index)
	}
	return string(str)
}

//随机生成一个日期类型数据
func MakeRandDate() string {
	year := MakeRandInt(1970, 2021)
	month := MakeRandInt(1, 12)
	var day int64
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = MakeRandInt(1, 31)
	case 4, 6, 9, 11:
		day = MakeRandInt(1, 30)
	case 2:
		day = MakeRandInt(1, 28)
	}

	strDate := strconv.FormatInt(year, 10) + "-" + strconv.FormatInt(month, 10) + "-" + strconv.FormatInt(day, 10)
	return strDate
}
