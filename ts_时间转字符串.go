package timezh

import (
	"time"
)

var TS = NewTs转字符串()

type ts转字符串 struct {
	D日期      *ts日期转字符串
	Y年份      *ts日期转字符串
	D月日      *ts日期转字符串
	T时刻      *ts时间转字符串
	T时间      *ts时间转字符串
	T日期和T和时刻 *ts时间转字符串
	T时间14位数  *ts时间转字符串
	D日期8位数   *ts日期转字符串
	T时刻6位数   *ts时间转字符串
}

func NewTs转字符串() *ts转字符串 {
	return &ts转字符串{
		D日期:      NewTs日期转字符串("2006-01-02"),
		Y年份:      NewTs日期转字符串("2006"),
		D月日:      NewTs日期转字符串("01-02"),
		T时刻:      NewTs时间转字符串("15:04:05"),
		T时间:      NewTs时间转字符串("2006-01-02 15:04:05"),
		T日期和T和时刻: NewTs时间转字符串("2006-01-02T15:04:05"),
		T时间14位数:  NewTs时间转字符串("20060102150405"),
		D日期8位数:   NewTs日期转字符串("20060102"),
		T时刻6位数:   NewTs时间转字符串("150405"),
	}
}

type ts日期转字符串 struct {
	s格式布局 string
}

func NewTs日期转字符串(s格式布局 string) *ts日期转字符串 {
	return &ts日期转字符串{
		s格式布局: s格式布局,
	}
}

func (ts *ts日期转字符串) Get转字符串(vTime time.Time) string {
	return vTime.Format(ts.s格式布局)
}

func (ts *ts日期转字符串) Get现在日期(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置))
}

func (ts *ts日期转字符串) Get当日日期(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置))
}

func (ts *ts日期转字符串) Get昨日日期(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置).AddDate(0, 0, -1))
}

func (ts *ts日期转字符串) Get最近含当日的前N个工作日(n int, z时区位置 *time.Location) (workdays []string) {
	return ts.Get某天含当日的前N个工作日(time.Now().In(z时区位置), n)
}

func (ts *ts日期转字符串) Get某天含当日的前N个工作日(vTime time.Time, n int) (workdays []string) {
	for {
		if WK.Is工作日(vTime) {
			workdays = append(workdays, ts.Get转字符串(vTime))
			if len(workdays) == n {
				return workdays
			}
		}
		vTime = vTime.AddDate(0, 0, -1)
	}
}

func (ts *ts日期转字符串) Get获取当月1号的日期(vTime time.Time) string {
	Y, M := TI.Get年份和月份(vTime)
	return IS.Get日期(Y, M, 1)
}

type ts时间转字符串 struct {
	s格式布局 string
}

func NewTs时间转字符串(s格式布局 string) *ts时间转字符串 {
	return &ts时间转字符串{
		s格式布局: s格式布局,
	}
}

func (ts *ts时间转字符串) Get转字符串(vTime time.Time) string {
	return vTime.Format(ts.s格式布局)
}

func (ts *ts时间转字符串) Get现在时间(z时区位置 *time.Location) string {
	return ts.Get转字符串(time.Now().In(z时区位置))
}
