package timezh

import "time"

var TI = &ti把时间转为数量{}

type ti把时间转为数量 struct{}

func (T *ti把时间转为数量) Get年份月份几号(t time.Time) (v年份, v月份, v几号 int) {
	v年份 = t.Year()
	v月份 = int(t.Month())
	v几号 = t.Day()
	return
}

func (T *ti把时间转为数量) Get几时几分秒数(t time.Time) (v几时, v几分, v秒数 int) {
	v几时 = t.Hour()
	v几分 = t.Minute()
	v秒数 = t.Second()
	return
}

func (T *ti把时间转为数量) Get数量(t time.Time) (v年份, v月份, v几号, v几时, v几分, v秒数 int) {
	v年份 = t.Year()
	v月份 = int(t.Month())
	v几号 = t.Day()
	v几时 = t.Hour()
	v几分 = t.Minute()
	v秒数 = t.Second()
	return
}

func (T *ti把时间转为数量) Get年份和月份(t time.Time) (v年份, v月份 int) {
	v年份 = t.Year()
	v月份 = int(t.Month())
	return
}

func (T *ti把时间转为数量) Get年份(t time.Time) (v年份 int) {
	v年份 = t.Year()
	return
}
