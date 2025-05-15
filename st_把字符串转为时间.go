package timezh

import (
	"time"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

var ST = &st把字符串转为时间{}

type st把字符串转为时间 struct{}

func (st *st把字符串转为时间) Get时间(s标准时间 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("2006-01-02 15:04:05", s标准时间, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return v时间
}

func (st *st把字符串转为时间) Get日期和分时(s标准时间没有秒数 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("2006-01-02 15:04", s标准时间没有秒数, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间没有秒数), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return v时间
}

func (st *st把字符串转为时间) Get通过日期得到时间(s标准日期 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("2006-01-02", s标准日期, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准日期), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return v时间
}

func (st *st把字符串转为时间) Get通过时刻得到当日时间(s标准时刻 string, z时区位置 *time.Location) time.Time {
	v时间, err := time.ParseInLocation("15:04:05", s标准时刻, z时区位置)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时刻), zap.Error(err))
		v时间 = time.Unix(0, 0)
	}
	return NT.New由当日时分秒得到时间(v时间.Hour(), v时间.Minute(), v时间.Second(), z时区位置)
}
