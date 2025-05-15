package timezh

import (
	"fmt"
	"testing"
	"time"
)

func Test_ts日期转字符串_Get最近含当日的前N个工作日(t *testing.T) {
	for _, x := range TS.D日期.Get最近含当日的前N个工作日(10, ZONE.CST) {
		fmt.Println(x)
	}
}

func Test_ts时间转字符串_Get现在时间(t *testing.T) {
	s := TS.T时间.Get现在时间(ZONE.CST)
	fmt.Println(s)
}

func Test_ts时间转字符串_Get转字符串(t *testing.T) {
	fmt.Println(TS.T时刻.Get转字符串(time.Now()))
}

func Test_ts日期转字符串_Get转字符串(t *testing.T) {
	fmt.Println(TS.D日期.Get转字符串(time.Now()))
}
