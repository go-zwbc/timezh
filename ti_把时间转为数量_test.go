package timezh

import (
	"fmt"
	"testing"
	"time"
)

func Test_ti把时间转为数量_Get几时几分秒数(t *testing.T) {
	fmt.Println(TI.Get几时几分秒数(time.Now().In(ZONE.CST)))
	fmt.Println(TI.Get几时几分秒数(time.Now().In(ZONE.ICT)))
}

func Test_ti把时间转为数量_Get年份月份几号(t *testing.T) {
	fmt.Println(TI.Get年份月份几号(time.Now().In(ZONE.CST)))
	fmt.Println(TI.Get年份月份几号(time.Now().In(ZONE.UTC)))
}

func Test_ti把时间转为数量_Get数量(t *testing.T) {
	fmt.Println(TI.Get数量(time.Now().In(ZONE.CST)))
	fmt.Println(TI.Get数量(time.Now().In(ZONE.UTC)))
}
