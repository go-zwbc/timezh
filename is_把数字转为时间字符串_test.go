package timezh_test

import (
	"testing"

	"github.com/go-zwbc/timezh"
	"github.com/stretchr/testify/require"
)

func Test_is把数字转为时间字符串_Get日期(t *testing.T) {
	res := timezh.IS.Get日期(2021, 1, 23)
	t.Log(res)
	require.Equal(t, "2021-01-23", res)
}
