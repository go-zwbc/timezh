package timezh

import (
	"fmt"
	"testing"
	"time"
)

func TestZone(t *testing.T) {
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 -0700"))
	fmt.Println(time.Now())
	fmt.Println(time.Now().In(ZONE.CST))
	{
		fmt.Println(time.Now().In(ZONE.UTC))
		fmt.Println(time.Now().In(ZONE.UTC).Zone())
	}
	{
		zoneTemp := time.FixedZone("自定义时区", 20)
		fmt.Println(time.Now().In(zoneTemp).Format("2006-01-02 15:04:05 -070000 MST"))
		fmt.Println(time.Now().In(zoneTemp).Zone())
	}
	fmt.Println(time.Now())
}

func TestFormat(t *testing.T) {
	fmt.Println(time.Now().In(ZONE.UTC).Format(time.RFC3339))
	fmt.Println(time.Now().In(ZONE.CST).Format(time.RFC1123))
}
