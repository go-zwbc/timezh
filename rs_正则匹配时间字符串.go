package timezh

import "regexp"

var RS = &rs正则匹配时间字符串{
	regexp日期:          regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
	regexp时刻:          regexp.MustCompile(`^\d{2}:\d{2}:\d{2}$`),
	regexp时间:          regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`),
	regexp日期和零时零分零秒时间: regexp.MustCompile(`^\d{4}-\d{2}-\d{2} 00:00:00$`),
	regexp日期8位数:       regexp.MustCompile(`^\d{8}$`),
	regexp时刻6位数:       regexp.MustCompile(`^\d{6}$`),
	regexp时刻8位带毫秒纯数:   regexp.MustCompile(`^\d{8}$`),
	regexp年份:          regexp.MustCompile(`^\d{4}$`),
}

type rs正则匹配时间字符串 struct {
	regexp日期          *regexp.Regexp
	regexp时刻          *regexp.Regexp
	regexp时间          *regexp.Regexp
	regexp日期和零时零分零秒时间 *regexp.Regexp
	regexp日期8位数       *regexp.Regexp
	regexp时刻6位数       *regexp.Regexp
	regexp时刻8位带毫秒纯数   *regexp.Regexp
	regexp年份          *regexp.Regexp
}

func (match *rs正则匹配时间字符串) Match日期(s string) bool {
	return match.regexp日期.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match时刻(s string) bool {
	return match.regexp时刻.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match时间(s string) bool {
	return match.regexp时间.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match日期和零时零分零秒时间(s string) bool {
	return match.regexp日期和零时零分零秒时间.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match日期8位数(s string) bool {
	return match.regexp日期8位数.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match时刻6位数(s string) bool {
	return match.regexp时刻6位数.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match时刻8位带毫秒纯数(s string) bool {
	return match.regexp时刻8位带毫秒纯数.MatchString(s)
}

func (match *rs正则匹配时间字符串) Match年份(s string) bool {
	return match.regexp年份.MatchString(s)
}
