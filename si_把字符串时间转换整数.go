package timezh

import (
	"fmt"
	"strconv"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

var SI = &si把字符串时间转换整数{}

type si把字符串时间转换整数 struct{}

func (si *si把字符串时间转换整数) Get通过日期得到年份(s标准日期 string) (v年份 int) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期))
		return 0
	}
	v年份, err := strconv.Atoi(s标准日期[0:4])
	if err != nil {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期), zap.Error(err))
		return 0
	}
	return v年份
}

func (si *si把字符串时间转换整数) Get通过日期得到年份月份哪日(s标准日期 string) (v年份, v月份, v几号 int) {
	num, err := fmt.Sscanf(s标准日期, "%4d-%2d-%2d", &v年份, &v月份, &v几号)
	if err != nil {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期), zap.Error(err))
		return 0, 0, 0
	}
	if num != 3 {
		zaplog.LOG.Warn("日期格式不对", zap.String("s", s标准日期), zap.Int("num", num))
		return 0, 0, 0
	}
	return v年份, v月份, v几号
}

func (si *si把字符串时间转换整数) Get通过时刻得到几时几分秒数(s标准时刻 string) (v几时, v几分, v秒数 int) {
	num, err := fmt.Sscanf(s标准时刻, "%2d:%2d:%2d", &v几时, &v几分, &v秒数)
	if err != nil {
		zaplog.LOG.Warn("时刻格式不对", zap.String("s", s标准时刻), zap.Error(err))
		return 0, 0, 0
	}
	if num != 3 {
		zaplog.LOG.Warn("时刻格式不对", zap.String("s", s标准时刻), zap.Int("num", num))
		return 0, 0, 0
	}
	return v几时, v几分, v秒数
}

func (si *si把字符串时间转换整数) Get通过时间得到整数(s标准时间 string) (v年份, v月份, v几号, v几时, v几分, v秒数 int) {
	if !RS.Match时间(s标准时间) {
		zaplog.LOG.Warn("时间格式不对", zap.String("s", s标准时间))
		return 0, 0, 0, 0, 0, 0
	}
	num, err := fmt.Sscanf(s标准时间, "%4d-%2d-%2d %2d:%2d:%2d", &v年份, &v月份, &v几号, &v几时, &v几分, &v秒数)
	if err != nil {
		zaplog.LOG.Warn("时间格式不对", zap.String("s", s标准时间), zap.Error(err))
		return 0, 0, 0, 0, 0, 0
	}
	if num != 6 {
		zaplog.LOG.Warn("时间格式不对", zap.String("s", s标准时间), zap.Int("num", num))
		return 0, 0, 0, 0, 0, 0
	}
	return v年份, v月份, v几号, v几时, v几分, v秒数
}
