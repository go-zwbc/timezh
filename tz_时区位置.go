package timezh

import "time"

var ZONE = struct {
	CST *time.Location //北京时间
	UTC *time.Location //世界标准时间
	GMT *time.Location //格林尼治标准时间 = UTC 时间
	ICT *time.Location //印度和中南半岛时间 柬埔寨时间
}{
	CST: time.FixedZone("CST", 8*3600), //北京时间
	UTC: time.UTC,                      //世界标准时间
	GMT: time.FixedZone("GMT", 0),      //格林尼治标准时间 = UTC
	ICT: time.FixedZone("ICT", 7*3600), //印度和中南半岛时间 柬埔寨时间
}
