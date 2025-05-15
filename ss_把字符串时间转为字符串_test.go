package timezh

import (
	"fmt"
	"testing"
)

func Test_ss把字符串时间转为字符串_Get把标准时刻转换为6位数时刻(t *testing.T) {
	HMS6n := SS.Get把标准时刻转换为6位数时刻("12:34:56")
	fmt.Println(HMS6n)
	HcMcS := SS.Get把6位数时刻转换为标准时刻(HMS6n)
	fmt.Println(HcMcS)
}
