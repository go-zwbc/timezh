package timezh

import (
	"time"
)

var WK = &wk时间和星期{}

type wk时间和星期 struct{}

func (wk *wk时间和星期) Is工作日(t time.Time) bool {
	week := t.Weekday()
	return week >= time.Monday && week <= time.Friday
}

func (wk *wk时间和星期) Is周六日(t time.Time) bool {
	week := t.Weekday()
	return week == time.Saturday || week == time.Sunday
}
