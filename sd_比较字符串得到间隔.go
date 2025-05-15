package timezh

import (
	"time"
)

var SD = &sd比较字符串得到间隔{}

type sd比较字符串得到间隔 struct{}

func (SD *sd比较字符串得到间隔) Get现在到日期的间隔(eYtMtD string, z时区位置 *time.Location) time.Duration {
	return SD.Get时间到日期的间隔(time.Now().In(z时区位置), eYtMtD)
}

func (SD *sd比较字符串得到间隔) Get时间到日期的间隔(t time.Time, eYtMtD string) time.Duration {
	return ST.Get通过日期得到时间(eYtMtD, t.Location()).Sub(t)
}

func (SD *sd比较字符串得到间隔) Get现在到当天时刻的间隔(eHcMcS string, z时区位置 *time.Location) time.Duration {
	return SD.Get时间到同天时刻的间隔(time.Now().In(z时区位置), eHcMcS)
}

func (SD *sd比较字符串得到间隔) Get时间到同天时刻的间隔(t time.Time, eHcMcS string) time.Duration {
	return ST.Get时间(TS.D日期.Get转字符串(t)+" "+eHcMcS, t.Location()).Sub(t)
}

func (SD *sd比较字符串得到间隔) Get现在到时间的间隔(eYtMtDbHcMcS string, z时区位置 *time.Location) time.Duration {
	return SD.Get时间到时间的间隔(time.Now().In(z时区位置), eYtMtDbHcMcS)
}

func (SD *sd比较字符串得到间隔) Get时间到时间的间隔(t time.Time, eYtMtDbHcMcS string) time.Duration {
	return ST.Get时间(eYtMtDbHcMcS, t.Location()).Sub(t)
}

func (SD *sd比较字符串得到间隔) Get两字符串时间间隔(sYtMtDbHcMcS, eYtMtDbHcMcS string) time.Duration {
	return ST.Get时间(eYtMtDbHcMcS, time.UTC).Sub(ST.Get时间(sYtMtDbHcMcS, time.UTC))
}
