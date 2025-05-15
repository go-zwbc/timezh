package timezh

import "time"

var NT = &nt把数字转化为时间{}

type nt把数字转化为时间 struct{}

func (nt *nt把数字转化为时间) New由当日时分秒得到时间(v几时, v几分, v几秒 int, z时区位置 *time.Location) time.Time {
	now := time.Now().In(z时区位置)
	return time.Date(now.Year(), now.Month(), now.Day(), v几时, v几分, v几秒, 0, z时区位置)
}

func (nt *nt把数字转化为时间) New由年月日得到零时时间(v年份, v月份, v几号 int, z时区位置 *time.Location) time.Time {
	return time.Date(v年份, time.Month(v月份), v几号, 0, 0, 0, 0, z时区位置)
}
