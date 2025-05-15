package timezh

import "time"

var TZ = struct {
	CST *time.Location //北京时间
	UTC *time.Location //世界标准时间
	GMT *time.Location //格林尼治标准时间 ~= UTC 时间
	ICT *time.Location //东南亚时间（泰国、越南、柬埔寨等）
}{
	CST: time.FixedZone("CST", 8*3600), //北京时间
	UTC: time.UTC,                      //世界标准时间
	GMT: time.FixedZone("GMT", 0),      //格林尼治标准时间 = UTC
	ICT: time.FixedZone("ICT", 7*3600), //东南亚时间（泰国、越南、柬埔寨等）
}
