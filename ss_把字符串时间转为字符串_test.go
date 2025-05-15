package timezh_test

import (
	"testing"

	"github.com/go-zwbc/timezh"
)

func Test_ss把字符串时间转为字符串_Get把标准时刻转换为6位数时刻(t *testing.T) {
	HMS6n := timezh.SS.Get把标准时刻转换为6位数时刻("12:34:56")
	t.Log(HMS6n)
	HcMcS := timezh.SS.Get把6位数时刻转换为标准时刻(HMS6n)
	t.Log(HcMcS)
}
