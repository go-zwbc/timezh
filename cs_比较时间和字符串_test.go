package timezh

import (
	"fmt"
	"testing"
	"time"
)

func Test_cs比较时间和字符串_I在时分秒内含边界(t *testing.T) {
	fmt.Println(CS.I在时分秒内含边界(time.Now().In(ZONE.CST), "09:15", "15:00"))
}
