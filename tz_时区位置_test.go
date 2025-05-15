package timezh_test

import (
	"testing"
	"time"

	"github.com/go-zwbc/timezh"
)

func TestTimeZone(t *testing.T) {
	t.Log(time.Now().Zone())
	t.Log(time.Now().Format("2006-01-02 15:04:05 -0700"))
	t.Log(time.Now())
	t.Log(time.Now().In(timezh.TZ.CST))
	{
		t.Log(time.Now().In(timezh.TZ.UTC))
		t.Log(time.Now().In(timezh.TZ.UTC).Zone())
	}
	{
		zoneTemp := time.FixedZone("自定义时区", 20)
		t.Log(time.Now().In(zoneTemp).Format("2006-01-02 15:04:05 -070000 MST"))
		t.Log(time.Now().In(zoneTemp).Zone())
	}
	t.Log(time.Now())
}

func TestFormat(t *testing.T) {
	t.Log(time.Now().In(timezh.TZ.UTC).Format(time.RFC3339))
	t.Log(time.Now().In(timezh.TZ.CST).Format(time.RFC1123))
}
