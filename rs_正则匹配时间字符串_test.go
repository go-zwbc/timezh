package timezh_test

import (
	"testing"

	"github.com/go-zwbc/timezh"
)

func Test_rs正则匹配时间字符串_Match时间(t *testing.T) {
	t.Log(timezh.RS.Match时间("2018-01-01 10:00:30"))
}

func Test_rs正则匹配时间字符串_Match时刻(t *testing.T) {
	t.Log(timezh.RS.Match时刻("09:00:01"))
}

func Test_rs正则匹配时间字符串_Match日期(t *testing.T) {
	t.Log(timezh.RS.Match日期("2018-01-01"))
	t.Log(timezh.RS.Match日期("2018-01-010"))
	t.Log(timezh.RS.Match日期("2018-010-01"))
	t.Log(timezh.RS.Match日期("02018-01-01"))
	t.Log(timezh.RS.Match日期("2018-01-01 2018-01-01"))
}
