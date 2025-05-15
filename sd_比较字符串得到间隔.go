package timezh

import (
	"time"
)

var SD = &sd比较字符串得到间隔{}

type sd比较字符串得到间隔 struct{}

func (SD *sd比较字符串得到间隔) Get现在到日期的间隔(s标准日期 string, z时区位置 *time.Location) time.Duration {
	return SD.Get时间到日期的间隔(time.Now().In(z时区位置), s标准日期)
}

func (SD *sd比较字符串得到间隔) Get时间到日期的间隔(t time.Time, s标准日期 string) time.Duration {
	return ST.Get通过日期得到时间(s标准日期, t.Location()).Sub(t)
}

func (SD *sd比较字符串得到间隔) Get现在到当天时刻的间隔(s标准时刻 string, z时区位置 *time.Location) time.Duration {
	return SD.Get时间到同天时刻的间隔(time.Now().In(z时区位置), s标准时刻)
}

func (SD *sd比较字符串得到间隔) Get时间到同天时刻的间隔(t time.Time, s标准时刻 string) time.Duration {
	return ST.Get时间(TS.D日期.Get转字符串(t)+" "+s标准时刻, t.Location()).Sub(t)
}

func (SD *sd比较字符串得到间隔) Get现在到时间的间隔(s标准时间 string, z时区位置 *time.Location) time.Duration {
	return SD.Get时间到时间的间隔(time.Now().In(z时区位置), s标准时间)
}

func (SD *sd比较字符串得到间隔) Get时间到时间的间隔(t time.Time, s标准时间 string) time.Duration {
	return ST.Get时间(s标准时间, t.Location()).Sub(t)
}

func (SD *sd比较字符串得到间隔) Get两字符串时间间隔(s开始时间, e结束时间 string) time.Duration {
	return ST.Get时间(e结束时间, time.UTC).Sub(ST.Get时间(s开始时间, time.UTC))
}
