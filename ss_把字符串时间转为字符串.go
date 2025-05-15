package timezh

import (
	"fmt"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ss把字符串时间转为字符串 struct{}

var SS = &ss把字符串时间转为字符串{}

func (*ss把字符串时间转为字符串) Get把年月日和时分秒分开返回(s标准时间 string) (v标准日期, v标准时间 string) {
	if !RS.Match时间(s标准时间) {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间))
		return "", ""
	}
	num, err := fmt.Sscanf(s标准时间, "%s %s", &v标准日期, &v标准时间)
	if err != nil {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间), zap.Error(err))
		return "", ""
	}
	if num != 2 {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间), zap.Int("num", num))
		return "", ""
	}
	return v标准日期, v标准时间
}

func (*ss把字符串时间转为字符串) Get把8位数日期转换为标准日期(s八位日期 string) (s标准日期 string) {
	if !RS.Match日期8位数(s八位日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s八位日期))
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", s八位日期[0:4], s八位日期[4:6], s八位日期[6:8])
}

func (*ss把字符串时间转为字符串) Get把标准日期转换为8位数日期(s标准日期 string) (s八位日期 string) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s标准日期))
		return ""
	}
	return fmt.Sprintf("%s%s%s", s标准日期[0:4], s标准日期[5:7], s标准日期[8:10])
}

func (*ss把字符串时间转为字符串) Get把标准时间转换为8位数日期(s标准时间 string) (s八位日期 string) {
	if !RS.Match时间(s标准时间) {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s标准时间))
		return ""
	}
	return fmt.Sprintf("%s%s%s", s标准时间[0:4], s标准时间[5:7], s标准时间[8:10])
}

func (*ss把字符串时间转为字符串) Get获取日期里的年份(s标准日期 string) (v年份 string) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s标准日期))
		return ""
	}
	return s标准日期[0:4]
}

func (*ss把字符串时间转为字符串) Get获取日期中的月份和杠和几号(s标准日期 string) (v月份和杠和几号 string) {
	if !RS.Match日期(s标准日期) {
		zaplog.LOG.Warn("日期格式错误", zap.String("s", s标准日期))
		return ""
	}
	return fmt.Sprintf("%s-%s", s标准日期[5:7], s标准日期[8:10])
}

func (*ss把字符串时间转为字符串) Get把6位数时刻转换为标准时刻(s六位时刻 string) (s标准时刻 string) {
	if !RS.Match时刻6位数(s六位时刻) {
		zaplog.LOG.Warn("时刻格式错误", zap.String("s", s六位时刻))
		return ""
	}
	return fmt.Sprintf("%s:%s:%s", s六位时刻[0:2], s六位时刻[2:4], s六位时刻[4:6])
}

func (*ss把字符串时间转为字符串) Get把8位数带毫秒的时刻转换为标准时刻(s八位带毫秒时刻 string) (s标准时刻 string) {
	if !RS.Match时刻8位带毫秒纯数(s八位带毫秒时刻) {
		zaplog.LOG.Warn("时刻格式错误", zap.String("s", s八位带毫秒时刻))
		return ""
	}
	return fmt.Sprintf("%s:%s:%s", s八位带毫秒时刻[0:2], s八位带毫秒时刻[2:4], s八位带毫秒时刻[4:6])
}

func (*ss把字符串时间转为字符串) Get把标准时刻转换为6位数时刻(s标准时刻 string) (s六位时刻 string) {
	if !RS.Match时刻(s标准时刻) {
		zaplog.LOG.Warn("时刻格式错误", zap.String("s", s标准时刻))
		return ""
	}
	return fmt.Sprintf("%s%s%s", s标准时刻[0:2], s标准时刻[3:5], s标准时刻[6:8])
}

func (*ss把字符串时间转为字符串) Get把4位时分转换为时和杠和分(s四位时分数 string) (v时间和冒号和分数 string) {
	return fmt.Sprintf("%s:%s", s四位时分数[0:2], s四位时分数[2:4])
}

func (*ss把字符串时间转为字符串) Get通过标准时间得到日期(s string) (v标准日期 string) {
	if !RS.Match时间(s) {
		zaplog.LOG.Warn("时间格式错误", zap.String("s", s))
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", s[0:4], s[5:7], s[8:10])
}
