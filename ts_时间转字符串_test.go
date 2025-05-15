package timezh_test

import (
	"testing"
	"time"

	"github.com/go-zwbc/timezh"
)

func Test_ts日期转字符串_Get最近含当日的前N个工作日(t *testing.T) {
	for _, x := range timezh.TS.D日期.Get最近含当日的前N个工作日(10, timezh.TZ.CST) {
		t.Log(x)
	}
}

func Test_ts时间转字符串_Get现在时间(t *testing.T) {
	s := timezh.TS.T时间.Get现在时间(timezh.TZ.CST)
	t.Log(s)
}

func Test_ts时间转字符串_Get转字符串(t *testing.T) {
	t.Log(timezh.TS.T时刻.Get转字符串(time.Now()))
}

func Test_ts日期转字符串_Get转字符串(t *testing.T) {
	t.Log(timezh.TS.D日期.Get转字符串(time.Now()))
}
