package timezh_test

import (
	"testing"

	"github.com/go-zwbc/timezh"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_si把字符串时间转换整数_Get通过日期得到年份(t *testing.T) {
	v年份 := timezh.SI.Get通过日期得到年份("2021-01-23")
	require.Equal(t, 2021, v年份)
	t.Log(v年份)
}

func Test_si把字符串时间转换整数_Get通过时刻得到几时几分秒数(t *testing.T) {
	V几时, v几分, v秒数 := timezh.SI.Get通过时刻得到几时几分秒数("16:44:02")
	assert.Equal(t, 16, V几时)
	assert.Equal(t, 44, v几分)
	assert.Equal(t, 2, v秒数)
	t.Log(V几时, v几分, v秒数)
}

func Test_si把字符串时间转换整数_Get通过日期得到年份月份哪日(t *testing.T) {
	v年份, v月份, v几号 := timezh.SI.Get通过日期得到年份月份哪日("2021-01-23")
	assert.Equal(t, 2021, v年份)
	assert.Equal(t, 1, v月份)
	assert.Equal(t, 23, v几号)
	t.Log(v年份, v月份, v几号)
}

func Test_si把字符串时间转换整数_Get通过时间得到整数(t *testing.T) {
	v年份, v月份, v几号, V几时, v几分, v秒数 := timezh.SI.Get通过时间得到整数("2021-01-23 16:47:03")
	assert.Equal(t, 2021, v年份)
	assert.Equal(t, 1, v月份)
	assert.Equal(t, 23, v几号)
	assert.Equal(t, 16, V几时)
	assert.Equal(t, 47, v几分)
	assert.Equal(t, 3, v秒数)
	t.Log(v年份, v月份, v几号, V几时, v几分, v秒数)
}

func Test_si把字符串时间转换整数_Get通过时间得到整数1(t *testing.T) {
	Y, month, D, H, minute, S := timezh.SI.Get通过时间得到整数("2018-01-02 03:04:05")
	if Y != 2018 || month != 1 || D != 2 || H != 3 || minute != 4 || S != 5 {
		t.Error(Y, month, D, H, minute, S)
	}
}

func Test_si把字符串时间转换整数_Get通过时刻得到几时几分秒数1(t *testing.T) {
	H, M, S := timezh.SI.Get通过时刻得到几时几分秒数("09:08:07")
	if H != 9 || M != 8 || S != 7 {
		t.Error(H, M, S)
		return
	}
}
