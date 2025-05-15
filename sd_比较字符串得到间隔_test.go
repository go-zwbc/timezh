package timezh

import (
	"fmt"
	"testing"
)

func Test_sd比较字符串得到间隔_Get现在到日期的间隔(t *testing.T) {
	fmt.Println(SD.Get现在到日期的间隔("2038-01-01", ZONE.CST))
	fmt.Println(SD.Get现在到日期的间隔("2099-12-31", ZONE.CST))
}

func Test_sd比较字符串得到间隔_Get现在到当天时刻的间隔(t *testing.T) {
	fmt.Println(SD.Get现在到当天时刻的间隔("00:00:00", ZONE.UTC))
	fmt.Println(SD.Get现在到当天时刻的间隔("00:30:00", ZONE.UTC))
	fmt.Println(SD.Get现在到当天时刻的间隔("09:28:00", ZONE.UTC))
	fmt.Println(SD.Get现在到当天时刻的间隔("23:59:59", ZONE.UTC))

	fmt.Println(SD.Get现在到当天时刻的间隔("00:00:00", ZONE.CST))
	fmt.Println(SD.Get现在到当天时刻的间隔("00:30:00", ZONE.CST))
	fmt.Println(SD.Get现在到当天时刻的间隔("09:28:00", ZONE.CST))
	fmt.Println(SD.Get现在到当天时刻的间隔("23:59:59", ZONE.CST))
}

func Test_sd比较字符串得到间隔_Get现在到时间的间隔(t *testing.T) {
	fmt.Println(SD.Get现在到时间的间隔("2021-01-23 17:21:00", ZONE.UTC))
	fmt.Println(SD.Get现在到时间的间隔("2021-01-23 17:21:00", ZONE.CST))
}

func Test_sd比较字符串得到间隔_Get两字符串时间间隔(t *testing.T) {
	fmt.Println(SD.Get两字符串时间间隔("2021-01-23 17:21:00", "2021-02-24 18:22:01"))
	fmt.Println(SD.Get两字符串时间间隔("2021-01-23 17:21:00", "2021-02-24 18:22:01"))
}
