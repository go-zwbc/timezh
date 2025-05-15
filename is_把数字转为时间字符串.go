package timezh

import "fmt"

var IS = is把数字转为时间字符串{}

type is把数字转为时间字符串 struct{}

func (si *is把数字转为时间字符串) Get日期(v年份, v月份, v几号 int) string {
	return fmt.Sprintf("%04d-%02d-%02d", v年份, v月份, v几号)
}

func (si *is把数字转为时间字符串) Get时刻(v几时, v几分, v几秒 int) string {
	return fmt.Sprintf("%02d:%02d:%02d", v几时, v几分, v几秒)
}

func (si *is把数字转为时间字符串) Get时间(v年份, v月份, v几号, v几时, v几分, v几秒 int) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", v年份, v月份, v几号, v几时, v几分, v几秒)
}
