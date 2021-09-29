package server

import (
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	birthday := Birthday{
		Id:        1,
		Name:      "测试人员",
		TimeType:  1,
		TimeText:  "9-30",
		SendEmail: "t1@abc.com;t2@abc.com",
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now()
	now.In(timeLocation)
	notifyEmail(birthday, now, timeLocation)
}
