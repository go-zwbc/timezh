package timezh

import (
	"time"
)

var WK = &wk时间和星期{}

type wk时间和星期 struct{}

func (wk *wk时间和星期) Is工作日(T time.Time) bool {
	week := T.Weekday()
	return week >= time.Monday && week <= time.Friday
}

func (wk *wk时间和星期) Is周六日(T time.Time) bool {
	week := T.Weekday()
	return week == time.Saturday || week == time.Sunday
}
