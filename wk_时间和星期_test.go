package timezh_test

import (
	"testing"
	"time"

	"github.com/go-zwbc/timezh"
)

func Test_wk时间和星期_Is工作日(t *testing.T) {
	t.Log(timezh.WK.Is工作日(time.Now().In(timezh.TZ.UTC)))
}

func TestIsWeekdays(t *testing.T) {
	now := time.Now()
	for v几号 := 0; v几号 < 10; v几号++ {
		date := now.AddDate(0, 0, v几号)
		t.Log(date.Weekday(), int(date.Weekday()))
	}
}
