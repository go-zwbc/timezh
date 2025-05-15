package timezh

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ns把时间戳转为时间_Get纳秒时间(t *testing.T) {
	n纳秒数 := time.Now().UnixNano()
	v时间 := NS.Get纳秒时间(n纳秒数)
	t.Log(v时间)
	require.Equal(t, n纳秒数, v时间.UnixNano())
}

func Test_ns把时间戳转为时间_Get纳秒时间_0(t *testing.T) {
	v时间 := NS.Get纳秒时间(0)
	t.Log(v时间)
}

func Test_ns把时间戳转为时间_Get微秒时间(t *testing.T) {
	n微秒数 := time.Now().UnixMicro()
	v时间 := NS.Get微秒时间(n微秒数)
	t.Log(v时间)
	require.Equal(t, n微秒数, v时间.UnixMicro())
}

func Test_ns把时间戳转为时间_Get毫秒时间(t *testing.T) {
	n毫秒数 := time.Now().UnixMilli()
	v时间 := NS.Get毫秒时间(n毫秒数)
	t.Log(v时间)
	require.Equal(t, n毫秒数, v时间.UnixMilli())
}

func Test_ns把时间戳转为时间_Get秒数时间(t *testing.T) {
	n总秒数 := time.Now().Unix()
	v时间 := NS.Get秒数时间(n总秒数)
	t.Log(v时间)
	require.Equal(t, n总秒数, v时间.Unix())
}
