package timezh_test

import (
	"testing"

	"github.com/go-zwbc/timezh"
)

func Test_nt把数字转化为时间_New由当日时分秒得到时间(t *testing.T) {
	v时间 := timezh.NT.New由当日时分秒得到时间(3, 3, 3, timezh.TZ.CST)
	t.Log(v时间)
}

func Test_nt把数字转化为时间_New由年月日得到零时时间(t *testing.T) {
	v时间 := timezh.NT.New由年月日得到零时时间(2025, 5, 15, timezh.TZ.CST)
	t.Log(v时间)
}
