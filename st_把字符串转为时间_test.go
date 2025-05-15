package timezh_test

import (
	"testing"

	"github.com/go-zwbc/timezh"
)

func Test_st把字符串转为时间_Get时间(t *testing.T) {
	{
		v时间 := timezh.ST.Get时间("2025-05-15 15:34:06", timezh.TZ.ICT)
		t.Log(v时间.UTC())
	}
	{
		v时间 := timezh.ST.Get时间("2025-05-15 15:34:06", timezh.TZ.CST)
		t.Log(v时间.UTC())
	}
	{
		v时间 := timezh.ST.Get时间("2025-05-15 15:34:06", timezh.TZ.UTC)
		t.Log(v时间.UTC())
	}
}

func Test_st把字符串转为时间_Get日期和分时(t *testing.T) {
	v时间 := timezh.ST.Get日期和分时("2025-05-15 09:30", timezh.TZ.CST)
	t.Log(v时间)
}

func Test_st把字符串转为时间_Get通过日期得到时间(t *testing.T) {
	v时间 := timezh.ST.Get通过日期得到时间("2025-05-15", timezh.TZ.CST)
	t.Log(v时间)
}

func Test_st把字符串转为时间_Get通过时刻得到当日时间(t *testing.T) {
	v时间 := timezh.ST.Get通过时刻得到当日时间("15:03:22", timezh.TZ.CST)
	t.Log(v时间)
}
