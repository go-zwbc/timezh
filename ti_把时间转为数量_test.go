package timezh_test

import (
	"testing"
	"time"

	"github.com/go-zwbc/timezh"
)

func Test_ti把时间转为数量_Get几时几分秒数(t *testing.T) {
	t.Log(timezh.TI.Get几时几分秒数(time.Now().In(timezh.TZ.CST)))
	t.Log(timezh.TI.Get几时几分秒数(time.Now().In(timezh.TZ.ICT)))
}

func Test_ti把时间转为数量_Get年份月份几号(t *testing.T) {
	t.Log(timezh.TI.Get年份月份几号(time.Now().In(timezh.TZ.CST)))
	t.Log(timezh.TI.Get年份月份几号(time.Now().In(timezh.TZ.UTC)))
}

func Test_ti把时间转为数量_Get数量(t *testing.T) {
	t.Log(timezh.TI.Get数量(time.Now().In(timezh.TZ.CST)))
	t.Log(timezh.TI.Get数量(time.Now().In(timezh.TZ.UTC)))
}
