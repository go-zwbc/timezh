package timezh

import (
	"fmt"
	"testing"
	"time"
)

func Test_wk时间和星期_Is工作日(t *testing.T) {
	fmt.Println(WK.Is工作日(time.Now().In(ZONE.UTC)))
}

func TestIsWeekdays(t *testing.T) {
	now := time.Now()
	for v几号 := 0; v几号 < 10; v几号++ {
		date := now.AddDate(0, 0, v几号)
		fmt.Println(date.Weekday(), int(date.Weekday()))
	}
}
