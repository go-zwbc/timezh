package timezh

import (
	"time"
)

var CS = &cs比较时间和字符串{}

type cs比较时间和字符串 struct{}

func (cs *cs比较时间和字符串) I在时分秒内含边界(t time.Time, s起始时刻, e结束时刻 string) bool {
	v := TS.T时刻.Get转字符串(t)
	return v >= s起始时刻 && v <= e结束时刻
}

func (cs *cs比较时间和字符串) A时间在时分秒前面(t time.Time, s时刻 string) bool {
	return TS.T时刻.Get转字符串(t) > s时刻
}

func (cs *cs比较时间和字符串) B时间在时分秒后面(t time.Time, s时刻 string) bool {
	return TS.T时刻.Get转字符串(t) < s时刻
}
